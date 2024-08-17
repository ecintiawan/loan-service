package investor

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
)

func TestNew(t *testing.T) {
	type args struct {
		client database.DB
	}
	tests := []struct {
		name string
		args args
		want repository.Investor
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
		want    *entity.Investor
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: func() *database.MockDB {
					row := database.NewMockPgxRow(
						[]string{
							"id",
							"identification_number",
							"name",
							"email",
							"status",
							"created_at",
							"updated_at",
						},
						[]interface{}{
							int64(1),
							"",
							"",
							"",
							constant.GeneralStatusActive,
							defaultDate,
							defaultDate,
						},
					)

					mock := database.NewMockDB(ctrl)
					mock.EXPECT().
						QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(row)

					return mock
				}(),
			},
			args: defaultArgs,
			want: &entity.Investor{
				ID:                   int64(1),
				IdentificationNumber: "",
				Name:                 "",
				Email:                "",
				Status:               constant.GeneralStatusActive,
				CreatedAt:            defaultDate,
				UpdatedAt:            defaultDate,
			},
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
