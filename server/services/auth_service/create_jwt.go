package auth_service

import (
	"os"
	"time"

	"github.com/james/nexus-server/types"
	"github.com/kataras/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateAccessToken(player types.Player) (string, error) {
	token, err := jwt.Sign(jwt.HS256, jwtKey, player, jwt.MaxAge(15*time.Minute))

	return string(token), err
}

func CreateRefreshToken(player types.Player) (string, error) {
	token, err := jwt.Sign(jwt.HS256, jwtKey, player, jwt.MaxAge(7*24*time.Hour))

	return string(token), err
}
