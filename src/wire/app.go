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

func NewApplication(u *utils.Utils, mqtt mqtt.MqttClient) *application {
	e := gin.New()
	e.Routes()
	e.Use()

	return &application{
		engine: gin.Default(),
		utils:  u,
	}
}

func (a *application) Run() error {
	return a.engine.Run("0.0.0.0:8080")
}
