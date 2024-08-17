package handler

import (
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/responsewrapper"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config            *config.Config
	echo              *echo.Echo
	healthHandler     *Health
	loanHandler       *Loan
	investmentHandler *Investment
}

func NewServer(
	config *config.Config,
	healthHandler *Health,
	loanHandler *Loan,
	investmentHandler *Investment,
) *Server {
	e := echo.New()

	s := &Server{
		config:            config,
		echo:              e,
		healthHandler:     healthHandler,
		loanHandler:       loanHandler,
		investmentHandler: investmentHandler,
	}
	e.HTTPErrorHandler = s.errorHandler

	return s
}

func (s *Server) InitRoutes() {
	s.echo.GET("/", s.healthHandler.HealthCheck)
	s.echo.GET("/ping", s.healthHandler.HealthCheck)
	s.echo.Static("/download", s.config.Vendor.Upload.Path)

	v1 := s.echo.Group("v1")

	// Loan
	v1.GET("/loan", s.loanHandler.HandleGet)
	v1.POST("/loan", s.loanHandler.HandleCreate)
	v1.PUT("/loan/:id", s.loanHandler.HandleProceed)

	// Investment
	v1.GET("/investment", s.investmentHandler.HandleGet)
	v1.POST("/investment", s.investmentHandler.HandleInvest)
}

func (s *Server) ListenAndServe() {
	s.echo.Logger.Fatal(s.echo.Start(":" + s.config.App.Port))
}

func (s *Server) errorHandler(err error, c echo.Context) {
	s.echo.DefaultHTTPErrorHandler(responsewrapper.ErrorHandler(err, c), c)
}
