package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/src/handler"
	"github.com/gsxhnd/go-api-template/src/middleware"
)

type Routes struct {
	Middleware middleware.Middlewarer
	RootHandle handler.RootHandler
}

func (r *Routes) Init(g *gin.Engine) {
	rootRoutes := g.Group("/")
	rootRoutes.GET("", r.RootHandle.Ping)

}

var RouteSet = wire.NewSet(
	middleware.NewMiddleware,
	handler.NewRootHandle,
	wire.Struct(new(Routes), "Middleware",
		"RootHandle"))
