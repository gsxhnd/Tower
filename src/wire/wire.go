//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"github.com/google/wire"
	"github.com/gsxhnd/tower/src/mqtt"
	"github.com/gsxhnd/tower/src/utils"
)

func InitApp(filePath *string) (*application, error) {
	wire.Build(NewApplication, utils.NewUtils, utils.UtilsSet, mqtt.NewMqttClient)
	return &application{}, nil
}
