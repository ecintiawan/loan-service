package handler

import (
	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/responsewrapper"
	"github.com/labstack/echo/v4"
)

// Investment is a handler for http request related to Investment
type Investment struct {
	service service.Investment
}

// NewInvestment returns new Investment handler.
func NewInvestment(service service.Investment) *Investment {
	return &Investment{
		service: service,
	}
}

// HandleGet handles the http request process of getting investment
func (l *Investment) HandleGet(c echo.Context) error {
	var (
		ctx = c.Request().Context()

		result entity.InvestmentResult
		err    error
	)

	filter := transformToInvestmentFilter(c)
	result, err = l.service.Get(ctx, filter)
	if err != nil {
		return err
	}

	return responsewrapper.OK(c, constant.MessageSuccessGet, result)
}

// HandleInvest handles the http request process of creating investment
func (l *Investment) HandleInvest(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		model = &entity.Investment{}

		err error
	)

	err = c.Bind(model)
	if err != nil {
		return errorwrapper.E("invalid parameter", errorwrapper.CodeInvalid)
	}

	err = l.service.Invest(ctx, model)
	if err != nil {
		return err
	}

	return responsewrapper.Created(c, constant.MessageSuccessCreate, nil)
}
