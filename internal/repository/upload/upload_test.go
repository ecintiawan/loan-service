package upload

import (
	"context"
	"reflect"
	"testing"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/file"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		config      *config.Config
		fileManager file.File
	}
	tests := []struct {
		name string
		args args
		want repository.Upload
	}{
		{
			name: "success",
			args: args{},
			want: &repoImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config, tt.args.fileManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repoImpl_Upload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		config      *config.Config
		fileManager file.File
	}
	type args struct {
		ctx   context.Context
		model *entity.File
	}
	defaultArgs := args{
		ctx: context.Background(),
		model: &entity.File{
			FileName: "test.png",
		},
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
				config: &config.Config{
					Vendor: config.Vendor{
						Upload: config.UploadConfig{
							URL: "http://127.0.0.1:8080/download/%s",
						},
					},
				},
				fileManager: func() *file.MockFile {
					mock := file.NewMockFile(ctrl)
					mock.EXPECT().
						Write(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
			args: defaultArgs,
			want: "http://127.0.0.1:8080/download/test.png",
		},
		{
			name: "error on write",
			fields: fields{
				config: &config.Config{
					Vendor: config.Vendor{
						Upload: config.UploadConfig{
							URL: "http://127.0.0.1:8080/download/%s",
						},
					},
				},
				fileManager: func() *file.MockFile {
					mock := file.NewMockFile(ctrl)
					mock.EXPECT().
						Write(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoImpl{
				config:      tt.fields.config,
				fileManager: tt.fields.fileManager,
			}
			got, err := r.Upload(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repoImpl.Upload() = %v, want %v", got, tt.want)
			}
		})
	}
}
