package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/go-api-template/src/routes"
)

type application struct {
	engine *gin.Engine
}

func NewApplication(g *gin.Engine, r *routes.Routes) *application {
	r.Init(g)
	return &application{
		engine: g,
	}
}

func (a *application) Run() error {
	return a.engine.Run("0.0.0.0:8080")
}
