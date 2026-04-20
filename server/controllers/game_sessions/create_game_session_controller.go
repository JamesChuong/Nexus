package game_sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/services/game_session_service"
	"github.com/james/nexus-server/types"
)

func CreateGameSessionController(c *gin.Context) {
	var player types.Player

	err := c.ShouldBindJSON(&player)

	gameSession, err := game_session_service.CreateGameSessionService(player)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"game_session": gameSession, "message": "Game Session Created"})

}
