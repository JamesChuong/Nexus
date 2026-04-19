package game_session_service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/james/nexus-server/services/redis_service"
	"github.com/james/nexus-server/types"
)

func findExistingGameSessions(player types.Player) error {
	if player.SessionId == "" {
		return nil
	}

	gameSession, err := redis_service.Search[types.GameSession](
		"idx:game_session",
		"@sessionId:{"+player.SessionId+"}",
		"@status:{"+types.StatusActive.String()+"}",
	)
	if err != nil {
		if errors.Is(err, redis_service.ErrNoResult) {
			return nil
		}

		return err
	}

	if gameSession != nil {
		return errors.New("Player is already in an active game session")
	}

	return nil

}

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

	if err := redis_service.SetRedisObject[types.GameSession](sessionRedisKey, gameSession, &gameSession); err != nil {
		return types.GameSession{}, err
	}

	return gameSession, nil

}
