package state

import (
	"context"
	"reflect"
	"testing"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/service"
)

func TestDetermineState(t *testing.T) {
	type args struct {
		ctx    context.Context
		status constant.LoanStatus
		action service.LoanAction
	}
	tests := []struct {
		name    string
		args    args
		want    service.LoanState
		wantErr bool
	}{
		{
			name: "success proposed",
			args: args{
				ctx:    context.Background(),
				status: constant.StatusProposed,
			},
			want: &ProposedState{},
		},
		{
			name: "success approved",
			args: args{
				ctx:    context.Background(),
				status: constant.StatusApproved,
			},
			want: &ApprovedState{},
		},
		{
			name: "success invested",
			args: args{
				ctx:    context.Background(),
				status: constant.StatusInvested,
			},
			want: &InvestedState{},
		},
		{
			name: "success disbursed",
			args: args{
				ctx:    context.Background(),
				status: constant.StatusDisbursed,
			},
			want: &DisbursedState{},
		},
		{
			name: "unknown state",
			args: args{
				ctx:    context.Background(),
				status: 10,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetermineState(tt.args.ctx, tt.args.status, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetermineState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetermineState() = %v, want %v", got, tt.want)
			}
		})
	}
}
