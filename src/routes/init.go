package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gsxhnd/go-api-template/src/handler"
	"github.com/gsxhnd/go-api-template/src/middleware"
)

type Routes struct {
	Engine     *gin.Engine
	Middleware middleware.Middlewarer
	RootHandle handler.RootHandler
	DemoHandle handler.DemoHandler
}

func (r *Routes) Init() {
	r.Engine.Use(r.Middleware.RequestLog())

	rootRoutes := r.Engine.Group("/")
	rootRoutes.GET("", r.RootHandle.Ping)

	r.NewDemoRoute()
}

var RouteSet = wire.NewSet(
	gin.New,
	middleware.NewMiddleware,
	handler.NewRootHandle,
	handler.NewDemoHandle,

	wire.Struct(new(Routes),
		"Engine",
		"Middleware",
		"RootHandle",
		"DemoHandle",
	),
)
