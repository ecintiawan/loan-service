package loan

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/internal/service/loan/state"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
)

type LoanImpl struct {
	repo   repository.Loan
	action service.LoanAction
}

func NewLoanImpl(
	repo repository.Loan,
	action service.LoanAction,
) service.Loan {
	return &LoanImpl{
		repo:   repo,
		action: action,
	}
}

// Get will return loan data based on filter
func (l *LoanImpl) Get(
	ctx context.Context,
	filter *entity.LoanFilter,
) (entity.LoanResult, error) {
	filter.Validate()

	return l.repo.Get(ctx, filter)
}

// Create will insert initial loan data
func (l *LoanImpl) Create(
	ctx context.Context,
	model *entity.Loan,
) error {
	if !model.IsValid() {
		return errorwrapper.E("invalid parameter values", errorwrapper.CodeInvalid)
	}
	model.Status = constant.StatusProposed

	return l.repo.Create(ctx, model)
}

// Proceed is an action to go through all loan states for certain loan data
func (l *LoanImpl) Proceed(
	ctx context.Context,
	req *entity.LoanProceed,
) error {
	if !req.IsValid() {
		return errorwrapper.E("invalid parameter values", errorwrapper.CodeInvalid)
	}

	existing, err := l.repo.GetDetail(ctx, req.Data.ID)
	if err != nil {
		return err
	}
	req.Data.Status = existing.Status
	req.Data.Amount = existing.Amount
	req.Data.Rate = existing.Rate

	state, err := state.DetermineState(ctx, existing.Status, l.action)
	if err != nil {
		return err
	}

	switch req.Action {
	case constant.ActionApprove:
		return state.Approve(ctx, req)
	case constant.ActionInvest:
		return state.Invest(ctx, req)
	case constant.ActionDisburse:
		return state.Disburse(ctx, req)
	}

	return errorwrapper.E("invalid action", errorwrapper.CodeInvalid)
}
