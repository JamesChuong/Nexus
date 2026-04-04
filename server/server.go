package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/james/nexus-server/services/redis"
)

func main() {
	err := redis.InitializeRedisIndexes()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = router.Run()
	if err != nil {
		return
	}
}
