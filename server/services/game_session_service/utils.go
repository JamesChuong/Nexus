package game_session_service

import (
	"errors"

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
