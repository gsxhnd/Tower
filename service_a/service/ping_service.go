package service

import (
	"github.com/gsxhnd/go-api-template/service_a/db"
	"github.com/gsxhnd/go-api-template/utils"
)

type PingService interface{}
type pingService struct {
	logger utils.Logger
	db     *db.Database
}

func NewPingService(l utils.Logger, db *db.Database) PingService {
	return &pingService{
		logger: l,
		db:     db,
	}
}
