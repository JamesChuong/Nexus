package players

import (
	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/types"
	"github.com/james/nexus-server/services/players"
	"time"
	"github.com/kataras/jwt"
)

github.com/james/nexus-server/types"

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

	c.JSON(200, gin.H{"player": player})

}
