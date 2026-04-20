package redis_service

import "github.com/goccy/go-json"

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
