package handler

import (
	"reflect"
	"testing"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewLoan(t *testing.T) {
	type args struct {
		service service.Loan
	}
	tests := []struct {
		name string
		args args
		want *Loan
	}{
		{
			name: "success",
			want: &Loan{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoan(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoan_HandleGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		service service.Loan
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.LoanResult{
							List: []*entity.Loan{
								{
									ID: 1,
								},
							},
							Pagination: entity.Pagination{
								Count: 1,
							},
						}, nil)

					return mock
				}(),
			},
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want: "{\"data\":{\"List\":[{\"id\":1,\"borrower_id\":0,\"amount\":0,\"rate\":0,\"status\":0,\"created_by\":0,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"approved_at\":\"0001-01-01T00:00:00Z\",\"invested_at\":\"0001-01-01T00:00:00Z\",\"disbursed_at\":\"0001-01-01T00:00:00Z\"}],\"count\":1,\"row\":0,\"page\":0},\"message\":\"Success get data\",\"status\":\"OK\"}\n",
		},
		{
			name: "error on get",
			fields: fields{
				service: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.LoanResult{}, assert.AnError)

					return mock
				}(),
			},
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loan{
				service: tt.fields.service,
			}
			if err := l.HandleGet(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Loan.HandleGet() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.c.(*mockEchoContext).getResponseBody()
			if string(got) != tt.want {
				t.Errorf("Loan.HandleGet() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestLoan_HandleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		service service.Loan
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want: "{\"data\":null,\"message\":\"Success create data\",\"status\":\"Created\"}\n",
		},
		{
			name: "error on create",
			fields: fields{
				service: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Create(gomock.Any(), gomock.Any()).
						Return(assert.AnError)

					return mock
				}(),
			},
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loan{
				service: tt.fields.service,
			}
			if err := l.HandleCreate(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Loan.HandleCreate() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.c.(*mockEchoContext).getResponseBody()
			if string(got) != tt.want {
				t.Errorf("Loan.HandleCreate() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestLoan_HandleProceed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		service service.Loan
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Proceed(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want: "{\"data\":null,\"message\":\"Success update data\",\"status\":\"OK\"}\n",
		},
		{
			name: "error on proceed",
			fields: fields{
				service: func() *service.MockLoan {
					mock := service.NewMockLoan(ctrl)
					mock.EXPECT().
						Proceed(gomock.Any(), gomock.Any()).
						Return(assert.AnError)

					return mock
				}(),
			},
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Loan{
				service: tt.fields.service,
			}
			if err := l.HandleProceed(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Loan.HandleProceed() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.c.(*mockEchoContext).getResponseBody()
			if string(got) != tt.want {
				t.Errorf("Loan.HandleProceed() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
