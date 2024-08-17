package service

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
)

//go:generate mockgen -source=internal/service/service.go -destination=internal/service/service_mock.go -package=service

// Loan encapsulates loan related logics
type Loan interface {
	// Get will return loan data based on filter
	Get(
		ctx context.Context,
		filter *entity.LoanFilter,
	) (entity.LoanResult, error)

	// Create will insert initial loan data
	Create(
		ctx context.Context,
		model *entity.Loan,
	) error

	// Proceed is an action to go through all loan states for certain loan data
	Proceed(
		ctx context.Context,
		req *entity.LoanProceed,
	) error
}

// LoanState encapsulates loan state related logics
type LoanState interface {
	// Approve will proceed loan to approved state
	Approve(
		ctx context.Context,
		req *entity.LoanProceed,
	) error

	// Invest will proceed loan to invested state
	Invest(
		ctx context.Context,
		req *entity.LoanProceed,
	) error

	// Disburse will proceed loan to disbursed state
	Disburse(
		ctx context.Context,
		req *entity.LoanProceed,
	) error
}

// LoanAction encapsulates loan action related logics
type LoanAction interface {
	// Approve will proceed loan to approved state
	Approve(
		ctx context.Context,
		req *entity.LoanProceed,
	) error

	// Invest will proceed loan to invested state
	Invest(
		ctx context.Context,
		req *entity.LoanProceed,
	) error

	// Disburse will proceed loan to disbursed state
	Disburse(
		ctx context.Context,
		req *entity.LoanProceed,
	) error
}

// Investment encapsulates investment related logics
type Investment interface {
	// Get will return investment data based on filter
	Get(
		ctx context.Context,
		filter *entity.InvestmentFilter,
	) (entity.InvestmentResult, error)

	// Invest will insert initial investment data
	Invest(
		ctx context.Context,
		req *entity.Investment,
	) error
}

type Services struct {
	Loan
	Investment
}
