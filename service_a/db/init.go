package db

import (
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/utils"
)

type Database struct {
}

func NewDatabase(cfg *utils.Config, l utils.Logger) (*Database, error) {
	return &Database{}, nil
}

var DBSet = wire.NewSet(NewDatabase)
