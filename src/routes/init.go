package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/src/handler"
)

type Routes struct {
	TestHandle handler.TestHandler
}

func (r *Routes) Init(g *gin.Engine) {
	g.GET("/", r.TestHandle.HelloHandler)
}

var RouteSet = wire.NewSet(
	handler.NewTestHandle,
	wire.Struct(new(Routes),
		"TestHandle"))
