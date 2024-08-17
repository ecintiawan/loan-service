package state

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
)

// InvestedState is the third state in the cycle, invested is once total amount of invested is equal the loan principal
// 1. loan can have multiple investors, each with each their own amount
// 2. total of invested amount can not be bigger than the loan principal amount
// 3. once invested all investors will receive an email containing link to agreement letter (pdf)
type InvestedState struct {
	action service.LoanAction
}

func NewInvestedState(action service.LoanAction) *InvestedState {
	return &InvestedState{
		action: action,
	}
}

func (s *InvestedState) Approve(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}

func (s *InvestedState) Invest(ctx context.Context, req *entity.LoanProceed) error {
	return errorIneligibleAction
}

func (s *InvestedState) Disburse(ctx context.Context, req *entity.LoanProceed) error {
	return s.action.Disburse(ctx, req)
}
