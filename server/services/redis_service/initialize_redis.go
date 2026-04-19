package redis_service

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: os.Getenv("REDIS_PASSWORD"),
	DB:       0,
	Protocol: 2,
})

var ctx = context.Background()

func createIndex(indexName string, options *redis.FTCreateOptions, schema []*redis.FieldSchema) error {
	err := RedisClient.FTInfo(ctx, indexName).Err()
	if err != nil {
		// Clear existing data under the index if it already exists
		_, err = RedisClient.Do(ctx, "FT.DROPINDEX", indexName, "DD").Result()
	}

	err = RedisClient.FTCreate(ctx, indexName, options, schema...).Err()

	return err
}

func InitializeRedisIndexes() error {
	playerIndex := "idx:player"

	playerIndexOptions := &redis.FTCreateOptions{
		OnJSON: false,
		Prefix: []interface{}{"player:"},
	}

	playerSchema := []*redis.FieldSchema{
		{FieldName: "playerId", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "playerName", FieldType: redis.SearchFieldTypeText},
		{FieldName: "status", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "ipAddress", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "sessionId", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "lastPing", FieldType: redis.SearchFieldTypeTag},
	}

	err := createIndex(playerIndex, playerIndexOptions, playerSchema)

	if err != nil {
		return err
	}

	gameSessionIndex := "idx:game_session"

	gameSessionIndexOptions := &redis.FTCreateOptions{
		OnJSON: false,
		Prefix: []interface{}{"game_session:"},
	}

	gameSessionSchema := []*redis.FieldSchema{
		{FieldName: "hostId", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "status", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "transport", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "sessionId", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "maxPlayers", FieldType: redis.SearchFieldTypeNumeric},
		{FieldName: "createdAt", FieldType: redis.SearchFieldTypeTag},
		{FieldName: "updatedAt", FieldType: redis.SearchFieldTypeTag},
	}

	err = createIndex(gameSessionIndex, gameSessionIndexOptions, gameSessionSchema)

	return err
}
