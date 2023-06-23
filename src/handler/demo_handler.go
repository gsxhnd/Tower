package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/go-api-template/src/service"
	"github.com/gsxhnd/go-api-template/src/utils"
)

type DemoHandler interface {
	Ping(ctx *gin.Context)
}
type demoHandle struct {
	logger utils.Logger
	svc    service.DemoService
}

func NewDemoHandle(l utils.Logger, s service.DemoService) DemoHandler {
	return &demoHandle{
		logger: l,
		svc:    s,
	}
}

func (h *demoHandle) Ping(ctx *gin.Context) {
	ctx.JSON(200, "pong")
}
