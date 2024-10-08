// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http

import (
	"github.com/ecintiawan/loan-service/internal/app/http/handler"
	"github.com/ecintiawan/loan-service/internal/repository/investment"
	"github.com/ecintiawan/loan-service/internal/repository/investor"
	"github.com/ecintiawan/loan-service/internal/repository/loan"
	"github.com/ecintiawan/loan-service/internal/repository/notifier"
	"github.com/ecintiawan/loan-service/internal/repository/upload"
	investment2 "github.com/ecintiawan/loan-service/internal/service/investment"
	loan2 "github.com/ecintiawan/loan-service/internal/service/loan"
	"github.com/ecintiawan/loan-service/internal/service/loan/action"
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/database"
	"github.com/ecintiawan/loan-service/pkg/email"
	"github.com/ecintiawan/loan-service/pkg/file"
	"github.com/ecintiawan/loan-service/pkg/lock"
)

// Injectors from wire.go:

func InitHttp() *handler.Server {
	configConfig := config.NewConfig()
	health := handler.NewHealth()
	db := database.NewDB(configConfig)
	repositoryLoan := loan.New(db)
	repositoryInvestment := investment.New(db)
	repositoryInvestor := investor.New(db)
	fileFile := file.NewFileImpl()
	repositoryUpload := upload.New(configConfig, fileFile)
	emailEmail := email.NewEmailImpl(configConfig)
	repositoryNotifier := notifier.New(emailEmail)
	pdfGenerator := file.NewPDFGeneratorImpl()
	loanAction := action.NewLoanActionImpl(configConfig, repositoryLoan, repositoryInvestment, repositoryInvestor, repositoryUpload, repositoryNotifier, pdfGenerator)
	serviceLoan := loan2.NewLoanImpl(repositoryLoan, loanAction)
	handlerLoan := handler.NewLoan(serviceLoan)
	lockLock := lock.NewLockImpl()
	serviceInvestment := investment2.NewInvestmentImpl(repositoryInvestment, repositoryLoan, serviceLoan, lockLock)
	handlerInvestment := handler.NewInvestment(serviceInvestment)
	server := handler.NewServer(configConfig, health, handlerLoan, handlerInvestment)
	return server
}
