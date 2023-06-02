package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/go-api-template/src/routes"
	"github.com/gsxhnd/go-api-template/src/utils"
)

type application struct {
	engine *gin.Engine
	utils  *utils.Utils
}

func NewApplication(g *gin.Engine, r *routes.Routes, u *utils.Utils) *application {
	r.Init(g)
	return &application{
		engine: g,
		utils:  u,
	}
}

func (a *application) Run() error {
	return a.engine.Run("0.0.0.0:8080")
}
