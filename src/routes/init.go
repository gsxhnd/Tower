package routes

import "github.com/gin-gonic/gin"

type Routes struct{}

func NewRoutes(g *gin.Engine) *Routes {
	return &Routes{}
}
