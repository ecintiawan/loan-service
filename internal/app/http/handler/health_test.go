package handler

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewHealth(t *testing.T) {
	tests := []struct {
		name string
		want *Health
	}{
		{
			name: "success",
			want: &Health{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealth_HealthCheck(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				c: newMockEchoContext(&mockEchoContext{}),
			},
			want:    "pong",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Health{}
			if err := h.HealthCheck(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Health.HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.c.(*mockEchoContext).getResponseBody()
			if string(got) != tt.want {
				t.Errorf("Health.HealthCheck() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
