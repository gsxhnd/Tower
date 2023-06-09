package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/go-api-template/src/utils"
)

type RootHandler interface {
	Ping(ctx *gin.Context)
}
type rootHandle struct {
	logger utils.Logger
}

func NewRootHandle(l utils.Logger) RootHandler {
	return &rootHandle{
		logger: l,
	}
}

func (h *rootHandle) Ping(ctx *gin.Context) {
	ctx.JSON(200, "pong")
}
