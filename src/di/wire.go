//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/src/dao"
	"github.com/gsxhnd/go-api-template/src/routes"
	"github.com/gsxhnd/go-api-template/src/service"
	"github.com/gsxhnd/go-api-template/src/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		NewApplication,
		routes.RouteSet,
		utils.UtilsSet,
		service.ServiceSet,
		dao.DaoSet,
	)
	return &Application{}, nil
}

func InitTask() (*Task, error) {
	wire.Build(
		NewTask,
	)
	return &Task{}, nil
}
