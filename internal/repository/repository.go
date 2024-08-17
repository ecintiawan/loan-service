package repository

import (
	"context"

	"github.com/ecintiawan/loan-service/internal/entity"
)

//go:generate mockgen -source=internal/repository/repository.go -destination=internal/repository/repository_mock.go -package=repository

// Loan encapsulates loan related logics
type Loan interface {
	// Get will return loan data based on filter
	Get(
		ctx context.Context,
		filter *entity.LoanFilter,
	) (entity.LoanResult, error)

	// GetDetail will return loan data based on filter
	GetDetail(
		ctx context.Context,
		id int64,
	) (*entity.Loan, error)

	// Create will insert initial loan data
	Create(
		ctx context.Context,
		model *entity.Loan,
	) error

	// Proceed is an action to go through all loan states for certain loan data
	Update(
		ctx context.Context,
		model *entity.Loan,
	) error
}

// Investment encapsulates investment related logics
type Investment interface {
	// Get will return investment data based on filter
	Get(
		ctx context.Context,
		filter *entity.InvestmentFilter,
	) (entity.InvestmentResult, error)

	// GetAmountSum will return investment amount sum data based on filter
	GetAmountSum(
		ctx context.Context,
		filter *entity.InvestmentFilter,
	) (float64, error)

	// Create will insert initial investment data
	Create(
		ctx context.Context,
		model *entity.Investment,
	) error
}

// Investor encapsulates investor related logics
type Investor interface {
	// GetDetail will return investor data based on filter
	GetDetail(
		ctx context.Context,
		id int64,
	) (*entity.Investor, error)
}

// Upload encapsulates upload related logics
type Upload interface {
	// Upload will upload files based on model
	Upload(
		ctx context.Context,
		model *entity.File,
	) (string, error)
}

// Notifier encapsulates notifier related logics
type Notifier interface {
	// Notify will notify related entities based on model
	Notify(
		ctx context.Context,
		model *entity.Notifier,
	) error
}
