package state

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
)

// ProposedState is the first state in the cycle, when a loan is created it will have proposed state
type ProposedState struct {
	action service.LoanAction
}

func NewProposedState(action service.LoanAction) *ProposedState {
	return &ProposedState{
		action: action,
	}
}

func (s *ProposedState) Approve(ctx context.Context, req *entity.LoanProceed) error {
	return s.action.Approve(ctx, req)
}

func (s *ProposedState) Invest(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}

func (s *ProposedState) Disburse(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}
