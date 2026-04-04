package redisClient

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

RedisClient := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password
	DB:       0,  // use default DB
	Protocol: 2,
})

ctx := context.Background()

func InitializeRedisIndexes() {

}