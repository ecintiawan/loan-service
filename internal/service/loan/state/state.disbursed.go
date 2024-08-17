package state

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
)

// DisbursedState is the fourth and final state in the cycle, disbursed is when is loan is given to borrower
// this state should be final, meaning it can't be any more forwarded or backwarded
type DisbursedState struct {
	action service.LoanAction
}

func NewDisbursedState(action service.LoanAction) *DisbursedState {
	return &DisbursedState{
		action: action,
	}
}

func (s *DisbursedState) Approve(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}

func (s *DisbursedState) Invest(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}

func (s *DisbursedState) Disburse(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}
