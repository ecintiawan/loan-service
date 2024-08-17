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

func TestNewInvestment(t *testing.T) {
	type args struct {
		service service.Investment
	}
	tests := []struct {
		name string
		args args
		want *Investment
	}{
		{
			name: "success",
			want: &Investment{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvestment(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvestment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvestment_HandleGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		service service.Investment
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
				service: func() *service.MockInvestment {
					mock := service.NewMockInvestment(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.InvestmentResult{
							List: []*entity.Investment{
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
			want: "{\"data\":{\"List\":[{\"id\":1,\"investor_id\":0,\"loan_id\":0,\"amount\":0,\"roi\":0,\"status\":0,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}],\"count\":1,\"row\":0,\"page\":0},\"message\":\"Success get data\",\"status\":\"OK\"}\n",
		},
		{
			name: "error on get",
			fields: fields{
				service: func() *service.MockInvestment {
					mock := service.NewMockInvestment(ctrl)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.InvestmentResult{}, assert.AnError)

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
			l := &Investment{
				service: tt.fields.service,
			}
			if err := l.HandleGet(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Investment.HandleGet() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.c.(*mockEchoContext).getResponseBody()
			if string(got) != tt.want {
				t.Errorf("Investment.HandleGet() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestInvestment_HandleInvest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		service service.Investment
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
				service: func() *service.MockInvestment {
					mock := service.NewMockInvestment(ctrl)
					mock.EXPECT().
						Invest(gomock.Any(), gomock.Any()).
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
			name: "error on invest",
			fields: fields{
				service: func() *service.MockInvestment {
					mock := service.NewMockInvestment(ctrl)
					mock.EXPECT().
						Invest(gomock.Any(), gomock.Any()).
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
			l := &Investment{
				service: tt.fields.service,
			}
			if err := l.HandleInvest(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Investment.HandleInvest() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.c.(*mockEchoContext).getResponseBody()
			if string(got) != tt.want {
				t.Errorf("Investment.HandleInvest() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
