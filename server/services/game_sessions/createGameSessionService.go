package game_sessions

import (
	"github.com/james/nexus-server/services/redis_service"
	"github.com/james/nexus-server/types"
)

var redisClient = redis_service.RedisClient

func CreateGameSessionService(player types.Player) (types.GameSession, error) {
}
