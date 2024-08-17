package handler

import (
	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/responsewrapper"
	"github.com/labstack/echo/v4"
)

// Loan is a handler for http request related to Loan
type Loan struct {
	service service.Loan
}

// NewLoan returns new Loan handler.
func NewLoan(service service.Loan) *Loan {
	return &Loan{
		service: service,
	}
}

// HandleGet handles the http request process of getting loan
func (l *Loan) HandleGet(c echo.Context) error {
	var (
		ctx = c.Request().Context()

		result entity.LoanResult
		err    error
	)

	filter := transformToLoanFilter(c)
	result, err = l.service.Get(ctx, filter)
	if err != nil {
		return err
	}

	return responsewrapper.OK(c, constant.MessageSuccessGet, result)
}

// HandleCreate handles the http request process of creating loan
func (l *Loan) HandleCreate(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		model = &entity.Loan{}

		err error
	)

	err = c.Bind(model)
	if err != nil {
		return errorwrapper.E("error binding request", errorwrapper.CodeInvalid)
	}

	err = l.service.Create(ctx, model)
	if err != nil {
		return err
	}

	return responsewrapper.Created(c, constant.MessageSuccessCreate, nil)
}

// HandleProceed handles the http request process of proceeding loan
func (l *Loan) HandleProceed(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = &entity.LoanProceed{}

		err error
	)

	req = transformToLoanProceed(c)
	err = l.service.Proceed(ctx, req)
	if err != nil {
		return err
	}

	return responsewrapper.OK(c, constant.MessageSuccessUpdate, nil)
}
