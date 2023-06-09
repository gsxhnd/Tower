//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/src/mqtt"
	"github.com/gsxhnd/go-api-template/src/routes"
	"github.com/gsxhnd/go-api-template/src/utils"
)

func InitApp(filePath *string) (*application, error) {
	wire.Build(NewApplication, utils.UtilsSet, mqtt.NewMqttClient, gin.New, routes.RouteSet)
	return &application{}, nil
}
