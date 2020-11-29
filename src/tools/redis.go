package tools

import (
	"context"
	"github.com/doter1995/user_center/src/config"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var RCtx = context.Background()

func InitRedis() *redis.Client {
	cfg := config.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Url,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	Rdb = rdb
	return rdb
}
