package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/tower/src/utils"
)

type application struct {
	engine *gin.Engine
	config *utils.Config
}

func NewApplication(cfg *utils.Config) *application {
	e := gin.New()
	e.Routes()
	e.Use()

	return &application{
		engine: gin.Default(),
		config: cfg,
	}
}

func (a *application) Run() error {
	return a.engine.Run("0.0.0.0:8080")
}
