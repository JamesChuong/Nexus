package redis_service

import (
	"errors"
	"strings"

	"github.com/goccy/go-json"
)

var ErrNoResult = errors.New("no result matches this query")

// Wrapper functions around the native go-redis functions to handle marshaling to different types
// The return type must be specified when called
// Ex: Search[Player]('1234', '@PlayerName:test', @Status:connected') will return a Player object with name 'test' and status 'connected'

func Search[T any](index string, searchParams ...string) (any, error) {

	query := strings.Join(searchParams, " ")

	result, err := RedisClient.FTSearch(ctx, index, query).Result()

	if err != nil {
		return nil, err
	}

	if result.Total == 0 || len(result.Docs) == 0 {
		return nil, ErrNoResult
	}

	document := result.Docs[0].Fields

	var returnType T

	data, err := mapRedisObjectToStruct[T](document, &returnType)

	return data, err

}

func mapRedisObjectToStruct[T any](object map[string]string, returnType *T) (T, error) {
	var zero T

	b, err := json.Marshal(object)
	if err != nil {
		return zero, err
	}

	if err := json.Unmarshal(b, returnType); err != nil {
		return zero, err
	}

	return *returnType, nil
}

func SetRedisObject[T any](key string, interfaceType *T) error {
	b, err := json.Marshal(interfaceType)
	if err != nil {
		return err
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	if err := RedisClient.HSet(ctx, key, fields).Err(); err != nil {
		return err
	}

	return nil
}
