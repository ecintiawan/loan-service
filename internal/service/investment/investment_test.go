package investment

import (
	"context"
	"reflect"
	"testing"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/lock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewInvestmentImpl(t *testing.T) {
	type args struct {
		repoInvestment repository.Investment
		repoLoan       repository.Loan
		serviceLoan    service.Loan
		lock           lock.Lock
	}
	tests := []struct {
		name string
		args args
		want service.Investment
	}{
		{
			name: "success",
			args: args{},
			want: &InvestmentImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvestmentImpl(tt.args.repoInvestment, tt.args.repoLoan, tt.args.serviceLoan, tt.args.lock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvestmentImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvestmentImpl_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repoInvestment repository.Investment
		repoLoan       repository.Loan
		serviceLoan    service.Loan
		lock           lock.Lock
	}
	type args struct {
		ctx    context.Context
		filter *entity.InvestmentFilter
	}
	defaultArgs := args{
		ctx: context.Background(),
		filter: &entity.InvestmentFilter{
			Status: 1,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.InvestmentResult
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.InvestmentResult{
							List: []*entity.Investment{
								{
									ID:     1,
									Amount: 10000,
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
			want: entity.InvestmentResult{
				List: []*entity.Investment{
					{
						ID:     1,
						Amount: 10000,
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
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.InvestmentResult{
							Pagination: entity.Pagination{
								Count: 1,
							},
						}, assert.AnError)

					return mock
				}(),
			},
			args: defaultArgs,
			want: entity.InvestmentResult{
				Pagination: entity.Pagination{
					Count: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InvestmentImpl{
				repoInvestment: tt.fields.repoInvestment,
				repoLoan:       tt.fields.repoLoan,
				serviceLoan:    tt.fields.serviceLoan,
				lock:           tt.fields.lock,
			}
			got, err := i.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvestmentImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvestmentImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvestmentImpl_Invest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repoInvestment repository.Investment
		repoLoan       repository.Loan
		serviceLoan    service.Loan
		lock           lock.Lock
	}
	type args struct {
		ctx context.Context
		req *entity.Investment
	}
	defaultArgs := args{
		ctx: context.Background(),
		req: &entity.Investment{
			InvestorID: 1,
			LoanID:     3,
			Amount:     200000,
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
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), gomock.Any()).
						Return(float64(1000000), nil)
					mock.EXPECT().
						Create(gomock.Any(), &entity.Investment{
							InvestorID: 1,
							LoanID:     3,
							Amount:     200000,
							ROI:        10,
							Status:     constant.GeneralStatusActive,
						}).
						Return(nil)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 1200000,
							Rate:   10,
							Status: constant.StatusApproved,
						}, nil)

					return mock
				}(),
				serviceLoan: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Proceed(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(getLockKey(3))
					mock.EXPECT().
						Unlock(getLockKey(3))

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name: "invalid request",
			fields: fields{
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(gomock.Any())
					mock.EXPECT().
						Unlock(gomock.Any())

					return mock
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &entity.Investment{},
			},
			wantErr: true,
		},
		{
			name: "error get loan detail",
			fields: fields{
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(nil, assert.AnError)

					return mock
				}(),
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(getLockKey(3))
					mock.EXPECT().
						Unlock(getLockKey(3))

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "invalid loan detail status",
			fields: fields{
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 1200000,
							Rate:   10,
							Status: constant.StatusProposed,
						}, nil)

					return mock
				}(),
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(getLockKey(3))
					mock.EXPECT().
						Unlock(getLockKey(3))

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error get amount sum",
			fields: fields{
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), gomock.Any()).
						Return(float64(0), assert.AnError)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 1200000,
							Rate:   10,
							Status: constant.StatusApproved,
						}, nil)

					return mock
				}(),
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(getLockKey(3))
					mock.EXPECT().
						Unlock(getLockKey(3))

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "investment amount exceeded",
			fields: fields{
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), gomock.Any()).
						Return(float64(1200000), nil)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 1200000,
							Rate:   10,
							Status: constant.StatusApproved,
						}, nil)

					return mock
				}(),
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(getLockKey(3))
					mock.EXPECT().
						Unlock(getLockKey(3))

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error create",
			fields: fields{
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), gomock.Any()).
						Return(float64(1000000), nil)
					mock.EXPECT().
						Create(gomock.Any(), gomock.Any()).
						Return(assert.AnError)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Loan{
							ID:     3,
							Amount: 1200000,
							Rate:   10,
							Status: constant.StatusApproved,
						}, nil)

					return mock
				}(),
				lock: func() *lock.MockLock {
					mock := lock.NewMockLock(ctrl)
					mock.EXPECT().
						Lock(getLockKey(3))
					mock.EXPECT().
						Unlock(getLockKey(3))

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InvestmentImpl{
				repoInvestment: tt.fields.repoInvestment,
				repoLoan:       tt.fields.repoLoan,
				serviceLoan:    tt.fields.serviceLoan,
				lock:           tt.fields.lock,
			}
			if err := i.Invest(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("InvestmentImpl.Invest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
