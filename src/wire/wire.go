//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"github.com/google/wire"
	"github.com/gsxhnd/tower/src/utils"
)

func InitApp() *application {
	wire.Build(NewApplication, utils.NewConfig)
	return &application{}
}
