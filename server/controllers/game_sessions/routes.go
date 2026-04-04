package game_sessions

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	session := router.Group("/session")
	{
		session.POST("/create", CreateGameSessionController)
	}
}
