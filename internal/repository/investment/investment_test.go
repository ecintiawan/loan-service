package investment

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
		want repository.Investment
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
		filter *entity.InvestmentFilter
	}
	defaultArgs := args{
		ctx:    context.Background(),
		filter: &entity.InvestmentFilter{},
	}
	defaultDate := time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local)
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
							"investor_id",
							"loan_id",
							"amount",
							"roi",
							"status",
							"created_at",
							"updated_at",
						},
						[][]interface{}{
							{
								int64(1),
								int64(1),
								int64(1),
								float64(10000),
								float64(11000),
								constant.GeneralStatusActive,
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
			want: entity.InvestmentResult{
				List: []*entity.Investment{
					{
						ID:         int64(1),
						InvestorID: int64(1),
						LoanID:     int64(1),
						Amount:     float64(10000),
						ROI:        float64(11000),
						Status:     constant.GeneralStatusActive,
						CreatedAt:  defaultDate,
						UpdatedAt:  defaultDate,
					},
				},
				Pagination: entity.Pagination{
					Count: 1,
				},
			},
		},
		{
			name: "error on select",
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
							"investor_id",
							"loan_id",
							"amount",
							"roi",
							"status",
							"created_at",
							"updated_at",
						},
						[][]interface{}{
							{
								int64(1),
								int64(1),
								int64(1),
								float64(10000),
								float64(11000),
								constant.GeneralStatusActive,
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
			want: entity.InvestmentResult{
				List: []*entity.Investment{},
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

func Test_repoImpl_GetAmountSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		client database.DB
	}
	type args struct {
		ctx    context.Context
		filter *entity.InvestmentFilter
	}
	defaultArgs := args{
		ctx:    context.Background(),
		filter: &entity.InvestmentFilter{},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: func() *database.MockDB {
					row := database.NewMockPgxRow(
						[]string{
							"sum",
						},
						[]interface{}{
							float64(10000),
						},
					)

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						QueryRow(gomock.Any(), gomock.Any()).
						Return(row)

					return mock
				}(),
			},
			args: defaultArgs,
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				client: tt.fields.client,
			}
			got, err := r.GetAmountSum(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.GetAmountSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repoImpl.GetAmountSum() = %v, want %v", got, tt.want)
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
		model *entity.Investment
	}
	defaultArgs := args{
		ctx:   context.Background(),
		model: &entity.Investment{},
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
