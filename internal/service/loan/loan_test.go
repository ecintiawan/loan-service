package loan

import (
	"context"
	"reflect"
	"testing"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewLoanImpl(t *testing.T) {
	type args struct {
		repo   repository.Loan
		action service.LoanAction
	}
	tests := []struct {
		name string
		args args
		want service.Loan
	}{
		{
			name: "success",
			args: args{},
			want: &LoanImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoanImpl(tt.args.repo, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoanImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanImpl_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   repository.Loan
		action service.LoanAction
	}
	type args struct {
		ctx    context.Context
		filter *entity.LoanFilter
	}
	defaultArgs := args{
		ctx: context.Background(),
		filter: &entity.LoanFilter{
			Status: 1,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.LoanResult
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.LoanResult{
							List: []*entity.Loan{
								{
									ID:     1,
									Amount: 1000000,
								},
							},
							Pagination: entity.Pagination{
								Count: 1,
							},
						}, nil)

					return mock
				}(),
			},
			args: defaultArgs,
			want: entity.LoanResult{
				List: []*entity.Loan{
					{
						ID:     1,
						Amount: 1000000,
					},
				},
				Pagination: entity.Pagination{
					Count: 1,
				},
			},
		},
		{
			name: "error on get",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.LoanResult{
							Pagination: entity.Pagination{
								Count: 1,
							},
						}, assert.AnError)

					return mock
				}(),
			},
			args: defaultArgs,
			want: entity.LoanResult{
				Pagination: entity.Pagination{
					Count: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoanImpl{
				repo:   tt.fields.repo,
				action: tt.fields.action,
			}
			got, err := l.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoanImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoanImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   repository.Loan
		action service.LoanAction
	}
	type args struct {
		ctx   context.Context
		model *entity.Loan
	}
	defaultArgs := args{
		ctx: context.Background(),
		model: &entity.Loan{
			BorrowerID: 1,
			Amount:     2000000,
			Rate:       10,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), &entity.Loan{
							BorrowerID: 1,
							Amount:     2000000,
							Rate:       10,
							Status:     constant.StatusProposed,
						}).
						Return(nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name:   "invalid param",
			fields: fields{},
			args: args{
				ctx:   context.Background(),
				model: &entity.Loan{},
			},
			wantErr: true,
		},
		{
			name: "error on create",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), &entity.Loan{
							BorrowerID: 1,
							Amount:     2000000,
							Rate:       10,
							Status:     constant.StatusProposed,
						}).
						Return(assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoanImpl{
				repo:   tt.fields.repo,
				action: tt.fields.action,
			}
			if err := l.Create(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("LoanImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoanImpl_Proceed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   repository.Loan
		action service.LoanAction
	}
	type args struct {
		ctx context.Context
		req *entity.LoanProceed
	}
	defaultArgs := args{
		ctx: context.Background(),
		req: &entity.LoanProceed{
			Action: constant.ActionApprove,
			Data: &entity.Loan{
				ID: 3,
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 2000000,
							Status: constant.StatusProposed,
						}, nil)

					return mock
				}(),
				action: func() *service.MockLoanAction {
					mock := service.NewMockLoanAction(ctrl)
					mock.EXPECT().
						Approve(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name:   "invalid param",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{},
			},
			wantErr: true,
		},
		{
			name: "error on get detail",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(nil, assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "unknown state",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 2000000,
						}, nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on action",
			fields: fields{
				repo: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 2000000,
							Status: constant.StatusProposed,
						}, nil)

					return mock
				}(),
				action: func() *service.MockLoanAction {
					mock := service.NewMockLoanAction(ctrl)
					mock.EXPECT().
						Approve(gomock.Any(), gomock.Any()).
						Return(assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoanImpl{
				repo:   tt.fields.repo,
				action: tt.fields.action,
			}
			if err := l.Proceed(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("LoanImpl.Proceed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
