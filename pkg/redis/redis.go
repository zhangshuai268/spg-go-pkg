package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Engine struct {
	Orm *redis.Client
}

func InitRedisEngine(addr, password string, db int) (*Engine, error) {
	res := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	_, err := res.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}
	return &Engine{
		Orm: res,
	}, nil
}
