package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/controllers/game_sessions"
	"github.com/james/nexus-server/services/redis_service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("No .env file found")
	}

	// Initialize Redis indexes for searching
	err = redis_service.InitializeRedisIndexes()

	if err != nil {
		log.Fatal("Failed to initialize redis: " + err.Error())
	}
}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		game_sessions.RegisterRoutes(api)
	}
	err := router.Run()
	if err != nil {
		return
	}
}
