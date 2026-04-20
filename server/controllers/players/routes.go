package players

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	player := router.Group("/player")
	{
		player.POST("/create", CreatePlayerController)
	}
}
