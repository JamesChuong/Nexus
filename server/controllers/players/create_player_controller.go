package players

import (
	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/services/auth_service"
	"github.com/james/nexus-server/services/players"
	"github.com/james/nexus-server/types"
)

func CreatePlayerController(c *gin.Context) {

	var req = types.PlayerRequestType{}

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	player, err := players.CreatePlayerService(req)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	token, err := auth_service.CreateJWT(player)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"player": player, "token": token})
}
