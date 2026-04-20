package players

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/james/nexus-server/services/redis_service"
	"github.com/james/nexus-server/types"
)

func CreatePlayerService(req types.PlayerRequestType) (types.Player, error) {
	playerName, ipAddress, err := validateCreatePlayerRequest(req)
	if err != nil {
		return types.Player{}, err
	}

	now := time.Now().Unix()

	player := types.Player{
		PlayerName: playerName,
		IPAddress:  ipAddress,
		PlayerId:   uuid.New().String(),
		LastPing:   now,
		SessionId:  "",
		Status:     types.StatusDisconnected.String(),
	}

	playerRedisKey := "player:" + player.PlayerId

	if err := redis_service.SetRedisObject[types.Player](playerRedisKey, &player); err != nil {
		return types.Player{}, fmt.Errorf("failed to connect to server %s: %w", player.PlayerId, err)
	}

	return player, nil
}
