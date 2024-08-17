package state

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
)

// ApprovedState is the second state in the cycle, approved is once the loan is approved by our staff
// 1. once approved it can not go back to proposed state
// 2. once approved loan is ready to be offered to investors/lender
type ApprovedState struct {
	action service.LoanAction
}

func NewApprovedState(action service.LoanAction) *ApprovedState {
	return &ApprovedState{
		action: action,
	}
}

func (s *ApprovedState) Approve(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}

func (s *ApprovedState) Invest(ctx context.Context, req *entity.LoanProceed) error {
	return s.action.Invest(ctx, req)
}

func (s *ApprovedState) Disburse(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}
