package http

import (
	"github.com/ecintiawan/loan-service/internal/app/http/handler"
	investmentRepo "github.com/ecintiawan/loan-service/internal/repository/investment"
	investorRepo "github.com/ecintiawan/loan-service/internal/repository/investor"
	loanRepo "github.com/ecintiawan/loan-service/internal/repository/loan"
	notifierRepo "github.com/ecintiawan/loan-service/internal/repository/notifier"
	uploadRepo "github.com/ecintiawan/loan-service/internal/repository/upload"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/internal/service/investment"
	"github.com/ecintiawan/loan-service/internal/service/loan"
	"github.com/ecintiawan/loan-service/internal/service/loan/action"
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/database"
	"github.com/ecintiawan/loan-service/pkg/email"
	"github.com/ecintiawan/loan-service/pkg/file"
	"github.com/ecintiawan/loan-service/pkg/lock"
	"github.com/google/wire"
)

var (
	httpSet = wire.NewSet(
		config.NewConfig,
		database.NewDB,
		email.NewEmailImpl,
		file.NewFileImpl,
		file.NewPDFGeneratorImpl,
		lock.NewLockImpl,
		repositorySet,
		serviceSet,
		handlerSet,
		wire.Struct(new(service.Services), "*"),
	)

	handlerSet = wire.NewSet(
		handler.NewHealth,
		handler.NewLoan,
		handler.NewInvestment,
		handler.NewServer,
	)

	serviceSet = wire.NewSet(
		action.NewLoanActionImpl,
		loan.NewLoanImpl,
		investment.NewInvestmentImpl,
	)

	repositorySet = wire.NewSet(
		loanRepo.New,
		investmentRepo.New,
		investorRepo.New,
		uploadRepo.New,
		notifierRepo.New,
	)
)
