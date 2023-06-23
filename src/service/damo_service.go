package service

import (
	"github.com/gsxhnd/go-api-template/src/dao"
	"github.com/gsxhnd/go-api-template/src/utils"
)

type DemoService interface{}
type demoService struct {
	logger utils.Logger
	db     dao.Database
}

func NewDemoService(l utils.Logger, db dao.Database) DemoService {
	return &demoService{
		logger: l,
		db:     db,
	}
}
