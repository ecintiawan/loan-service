package state

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
)

var (
	errorUnknownState     = errorwrapper.E("unknown state", errorwrapper.CodeInvalid)
	errorIneligibleAction = errorwrapper.E("ineligible action", errorwrapper.CodeInvalid)
)

func DetermineState(
	ctx context.Context,
	status constant.LoanStatus,
	action service.LoanAction,
) (service.LoanState, error) {
	var (
		state service.LoanState
		err   error
	)

	switch status {
	case constant.StatusProposed:
		// Initial state, after the loan has been created
		// May only do approval action
		state = NewProposedState(action)

	case constant.StatusApproved:
		// Second state, after the loan has been approved
		// May only do invest action after it's been fully invested
		state = NewApprovedState(action)

	case constant.StatusInvested:
		// Third state, after the loan has been invested
		// May only do disbursement action
		state = NewInvestedState(action)

	case constant.StatusDisbursed:
		// Fourth and final state, after the loan has been disbursed
		// No more action can be done
		state = NewDisbursedState(action)

	default:
		// Impossible state, throw error
		err = errorUnknownState
	}

	return state, err
}
