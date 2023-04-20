package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/tower/src/mqtt"
	"github.com/gsxhnd/tower/src/utils"
)

type application struct {
	engine *gin.Engine
	mqtt   mqtt.MqttClient
	utils  *utils.Utils
}

func NewApplication(g *gin.Engine, u *utils.Utils, mqtt mqtt.MqttClient) *application {
	return &application{
		engine: g,
		utils:  u,
	}
}

func (a *application) Run() error {
	return a.engine.Run("0.0.0.0:8080")
}
