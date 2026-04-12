package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/controllers/game_sessions"
	"github.com/james/nexus-server/services/redis_service"
)

func main() {
	err := redis_service.InitializeRedisIndexes()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	api := router.Group("/api")
	{
		game_sessions.RegisterRoutes(api)
	}
	err = router.Run()
	if err != nil {
		return
	}
}
