//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/middleware"
	"github.com/gsxhnd/go-api-template/service_a/db"
	"github.com/gsxhnd/go-api-template/service_a/handler"
	"github.com/gsxhnd/go-api-template/service_a/routes"
	"github.com/gsxhnd/go-api-template/service_a/service"
	"github.com/gsxhnd/go-api-template/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		utils.UtilsSet,
		NewApplication,
		routes.NewRouter,
		middleware.NewMiddleware,
		handler.HandlerSet,
		service.ServiceSet,
		db.DBSet,
	)
	return &Application{}, nil
}

func InitTask() (*Task, error) {
	wire.Build(
		NewTask,
	)
	return &Task{}, nil
}
