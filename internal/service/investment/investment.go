package investment

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/lock"
)

type InvestmentImpl struct {
	repoInvestment repository.Investment
	repoLoan       repository.Loan
	serviceLoan    service.Loan
	lock           lock.Lock
}

func NewInvestmentImpl(
	repoInvestment repository.Investment,
	repoLoan repository.Loan,
	serviceLoan service.Loan,
	lock lock.Lock,
) service.Investment {
	return &InvestmentImpl{
		repoInvestment: repoInvestment,
		repoLoan:       repoLoan,
		serviceLoan:    serviceLoan,
		lock:           lock,
	}
}

// Get will return investment data based on filter
func (i *InvestmentImpl) Get(
	ctx context.Context,
	filter *entity.InvestmentFilter,
) (entity.InvestmentResult, error) {
	filter.Validate()

	return i.repoInvestment.Get(ctx, filter)
}

// Invest will insert initial investment data
func (i *InvestmentImpl) Invest(
	ctx context.Context,
	req *entity.Investment,
) error {
	// locking to prevent racing investment case on the same loan data
	lockKey := getLockKey(req.LoanID)
	i.lock.Lock(lockKey)
	defer i.lock.Unlock(lockKey)

	if !req.IsValid() {
		return errorwrapper.E("invalid parameter values", errorwrapper.CodeInvalid)
	}

	// validate loan status
	loan, err := i.repoLoan.GetDetail(ctx, req.LoanID)
	if err != nil {
		return err
	}
	if loan.Status != constant.StatusApproved {
		return errorwrapper.E("loan status must be approved", errorwrapper.CodeInvalid)
	}

	// validate if the additional amount will exceed the principle amount
	amountSum, err := i.repoInvestment.GetAmountSum(ctx, &entity.InvestmentFilter{
		LoanID: req.LoanID,
		Status: constant.GeneralStatusActive,
	})
	if err != nil {
		return err
	}
	if req.Amount+amountSum > loan.Amount {
		return errorwrapper.E("investment amount exceeds remaining principle amount", errorwrapper.CodeInvalid)
	}

	req.ROI = loan.Rate
	req.Status = constant.GeneralStatusActive
	err = i.repoInvestment.Create(ctx, req)
	if err != nil {
		return err
	}

	// trigger loan to proceed to invested state
	if req.Amount+amountSum == loan.Amount {
		loan.InvestedAt = time.Now()
		errProceed := i.serviceLoan.Proceed(ctx, &entity.LoanProceed{
			Action: constant.ActionInvest,
			Data:   loan,
		})
		if errProceed != nil {
			log.Println("error proceeding loan to invested", err)
		}
	}

	return nil
}

func getLockKey(loanID int64) string {
	return fmt.Sprintf("investment:invest:%d", loanID)
}
