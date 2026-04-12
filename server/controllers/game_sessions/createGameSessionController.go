package game_sessions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/services/game_sessions"
	"github.com/james/nexus-server/types"
)

func CreateGameSessionController(c *gin.Context) {
	var player types.Player

	err := c.ShouldBindJSON(&player)

	game_session, err := game_sessions.CreateGameSessionService(player)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"game_session": game_session, "message": "Game Session Created"})

}
