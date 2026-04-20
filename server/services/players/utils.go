package players

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/james/nexus-server/types"
)

const maxPlayerNameLength = 32

func validateCreatePlayerRequest(req types.PlayerRequestType) (string, string, error) {
	playerName := strings.TrimSpace(req.PlayerName)
	if playerName == "" {
		return "", "", errors.New("player name is required")
	}
	if len(playerName) > maxPlayerNameLength {
		return "", "", fmt.Errorf("player name must be %d characters or fewer", maxPlayerNameLength)
	}

	ipAddress := strings.TrimSpace(req.IPAddress)
	if ipAddress == "" {
		return "", "", errors.New("ip address is required")
	}
	if net.ParseIP(ipAddress) == nil {
		return "", "", errors.New("ip address is invalid")
	}

	return playerName, ipAddress, nil
}
