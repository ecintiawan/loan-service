//go:build wireinject
// +build wireinject

package http

import (
	"github.com/ecintiawan/loan-service/internal/app/http/handler"
	"github.com/google/wire"
)

func InitHttp() *handler.Server {
	wire.Build(httpSet)
	return &handler.Server{}
}
