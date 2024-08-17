package notifier

import (
	"context"
	"reflect"
	"testing"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/email"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		emailer email.Email
	}
	tests := []struct {
		name string
		args args
		want repository.Notifier
	}{
		{
			name: "success",
			args: args{},
			want: &repoImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.emailer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoImpl_Notify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		emailer email.Email
	}
	type args struct {
		ctx   context.Context
		model *entity.Notifier
	}
	defaultArgs := args{
		ctx:   context.Background(),
		model: &entity.Notifier{},
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
				emailer: func() *email.MockEmail {
					mock := email.NewMockEmail(ctrl)
					mock.EXPECT().
						Send(gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name: "error on send",
			fields: fields{
				emailer: func() *email.MockEmail {
					mock := email.NewMockEmail(ctrl)
					mock.EXPECT().
						Send(gomock.Any()).
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
			r := &repoImpl{
				emailer: tt.fields.emailer,
			}
			if err := r.Notify(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
