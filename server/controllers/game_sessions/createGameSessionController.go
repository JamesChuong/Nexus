package game_sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/types"
)

func CreateGameSessionController(c *gin.Context) {
	var player types.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game Session Created"})

}
