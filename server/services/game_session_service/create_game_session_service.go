package game_session_service

import (
	"time"

	"github.com/google/uuid"
	"github.com/james/nexus-server/services/redis_service"
	"github.com/james/nexus-server/types"
)

func CreateGameSessionService(player types.Player) (types.GameSession, error) {

	err := findExistingGameSessions(player)
	if err != nil {
		return types.GameSession{}, err
	}

	gameSession := types.GameSession{
		HostId:     player.PlayerId,
		Status:     types.StatusActive.String(),
		Transport:  types.TransportP2P.String(),
		SessionId:  uuid.New().String(),
		MaxPlayers: 6,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	sessionRedisKey := "game_session:" + gameSession.SessionId

	if err := redis_service.SetRedisObject[types.GameSession](sessionRedisKey, &gameSession); err != nil {
		return types.GameSession{}, err
	}

	return gameSession, nil

}
