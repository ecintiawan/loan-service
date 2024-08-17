package loan

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/database"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		client database.DB
	}
	tests := []struct {
		name string
		args args
		want repository.Loan
	}{
		{
			name: "success",
			args: args{},
			want: &repoImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoImpl_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		client database.DB
	}
	type args struct {
		ctx    context.Context
		filter *entity.LoanFilter
	}
	defaultArgs := args{
		ctx:    context.Background(),
		filter: &entity.LoanFilter{},
	}
	defaultDate := time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local)
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
				client: func() *database.MockDB {
					row := database.NewMockPgxRow(
						[]string{
							"count",
						},
						[]interface{}{
							int64(1),
						},
					)

					rows := database.NewMockPgxRows(
						[]string{
							"id",
							"borrower_id",
							"amount",
							"rate",
							"approval_proof_url",
							"agreement_letter_url",
							"status",
							"created_by",
							"approved_by",
							"disbursed_by",
							"created_at",
							"updated_at",
							"approved_at",
							"invested_at",
							"disbursed_at",
						},
						[][]interface{}{
							{
								int64(1),
								int64(1),
								float64(10000),
								float64(10),
								"",
								"",
								constant.StatusProposed,
								int64(1),
								int64(1),
								int64(1),
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
							},
						},
					)

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						QueryRow(gomock.Any(), gomock.Any()).
						Return(row)
					mock.EXPECT().
						Query(gomock.Any(), gomock.Any()).
						Return(rows, nil)

					return mock
				}(),
			},
			args: defaultArgs,
			want: entity.LoanResult{
				List: []*entity.Loan{
					{
						ID:                 int64(1),
						BorrowerID:         int64(1),
						Amount:             float64(10000),
						Rate:               float64(10),
						ApprovalProofURL:   "",
						AgreementLetterURL: "",
						Status:             constant.StatusProposed,
						CreatedBy:          int64(1),
						ApprovedBy:         int64(1),
						DisbursedBy:        int64(1),
						CreatedAt:          defaultDate,
						UpdatedAt:          defaultDate,
						ApprovedAt:         defaultDate,
						InvestedAt:         defaultDate,
						DisbursedAt:        defaultDate,
					},
				},
				Pagination: entity.Pagination{
					Count: 1,
				},
			},
		},
		{
			name: "error select",
			fields: fields{
				client: func() *database.MockDB {
					row := database.NewMockPgxRow(
						[]string{
							"count",
						},
						[]interface{}{
							int64(1),
						},
					)

					rows := database.NewMockPgxRows(
						[]string{
							"id",
							"borrower_id",
							"amount",
							"rate",
							"approval_proof_url",
							"agreement_letter_url",
							"status",
							"created_by",
							"approved_by",
							"disbursed_by",
							"created_at",
							"updated_at",
							"approved_at",
							"invested_at",
							"disbursed_at",
						},
						[][]interface{}{
							{
								int64(1),
								int64(1),
								float64(10000),
								float64(10),
								"",
								"",
								constant.StatusProposed,
								int64(1),
								int64(1),
								int64(1),
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
							},
						},
					)

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						QueryRow(gomock.Any(), gomock.Any()).
						Return(row)
					mock.EXPECT().
						Query(gomock.Any(), gomock.Any()).
						Return(rows, assert.AnError)

					return mock
				}(),
			},
			args: defaultArgs,
			want: entity.LoanResult{
				List: []*entity.Loan{},
				Pagination: entity.Pagination{
					Count: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				client: tt.fields.client,
			}
			got, err := r.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repoImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoImpl_GetDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		client database.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	defaultArgs := args{
		ctx: context.Background(),
		id:  3,
	}
	defaultDate := time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Loan
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: func() *database.MockDB {
					rows := database.NewMockPgxRows(
						[]string{
							"id",
							"borrower_id",
							"amount",
							"rate",
							"approval_proof_url",
							"agreement_letter_url",
							"status",
							"created_by",
							"approved_by",
							"disbursed_by",
							"created_at",
							"updated_at",
							"approved_at",
							"invested_at",
							"disbursed_at",
						},
						[][]interface{}{
							{
								int64(1),
								int64(1),
								float64(10000),
								float64(10),
								"",
								"",
								constant.StatusProposed,
								int64(1),
								int64(1),
								int64(1),
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
							},
						},
					)

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Query(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(rows, nil)

					return mock
				}(),
			},
			args: defaultArgs,
			want: &entity.Loan{
				ID:                 int64(1),
				BorrowerID:         int64(1),
				Amount:             float64(10000),
				Rate:               float64(10),
				ApprovalProofURL:   "",
				AgreementLetterURL: "",
				Status:             constant.StatusProposed,
				CreatedBy:          int64(1),
				ApprovedBy:         int64(1),
				DisbursedBy:        int64(1),
				CreatedAt:          defaultDate,
				UpdatedAt:          defaultDate,
				ApprovedAt:         defaultDate,
				InvestedAt:         defaultDate,
				DisbursedAt:        defaultDate,
			},
		},
		{
			name: "error select",
			fields: fields{
				client: func() *database.MockDB {
					rows := database.NewMockPgxRows(
						[]string{
							"id",
							"borrower_id",
							"amount",
							"rate",
							"approval_proof_url",
							"agreement_letter_url",
							"status",
							"created_by",
							"approved_by",
							"disbursed_by",
							"created_at",
							"updated_at",
							"approved_at",
							"invested_at",
							"disbursed_at",
						},
						[][]interface{}{
							{
								int64(1),
								int64(1),
								float64(10000),
								float64(10),
								"",
								"",
								constant.StatusProposed,
								int64(1),
								int64(1),
								int64(1),
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
								defaultDate,
							},
						},
					)

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Query(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(rows, assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				client: tt.fields.client,
			}
			got, err := r.GetDetail(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.GetDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repoImpl.GetDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		client database.DB
	}
	type args struct {
		ctx   context.Context
		model *entity.Loan
	}
	defaultArgs := args{
		ctx:   context.Background(),
		model: &entity.Loan{},
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
				client: func() *database.MockDB {
					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(&database.MockPgxTx{}, nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name: "error on begin",
			fields: fields{
				client: func() *database.MockDB {
					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(nil, assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on exec",
			fields: fields{
				client: func() *database.MockDB {
					tx := &database.MockPgxTx{}
					tx.ExecFunc = func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
						return pgconn.CommandTag{}, assert.AnError
					}

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(tx, nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on commit",
			fields: fields{
				client: func() *database.MockDB {
					tx := &database.MockPgxTx{}
					tx.ExecFunc = func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
						return pgconn.CommandTag{}, nil
					}
					tx.CommitFunc = func(ctx context.Context) error {
						return assert.AnError
					}

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(tx, nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				client: tt.fields.client,
			}
			if err := r.Create(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repoImpl_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		client database.DB
	}
	type args struct {
		ctx   context.Context
		model *entity.Loan
	}
	defaultArgs := args{
		ctx:   context.Background(),
		model: &entity.Loan{},
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
				client: func() *database.MockDB {
					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(&database.MockPgxTx{}, nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name: "error on begin",
			fields: fields{
				client: func() *database.MockDB {
					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(nil, assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on exec",
			fields: fields{
				client: func() *database.MockDB {
					tx := &database.MockPgxTx{}
					tx.ExecFunc = func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
						return pgconn.CommandTag{}, assert.AnError
					}

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(tx, nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on commit",
			fields: fields{
				client: func() *database.MockDB {
					tx := &database.MockPgxTx{}
					tx.ExecFunc = func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
						return pgconn.CommandTag{}, nil
					}
					tx.CommitFunc = func(ctx context.Context) error {
						return assert.AnError
					}

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						Begin(gomock.Any()).
						Return(tx, nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				client: tt.fields.client,
			}
			if err := r.Update(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
