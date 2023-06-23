package dao

import (
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/src/utils"
)

type Database interface{}
type database struct{}

func NewDatabase(l utils.Logger) Database {
	return &database{}
}

var DaoSet = wire.NewSet(NewDatabase)
