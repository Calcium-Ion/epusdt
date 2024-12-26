package dao

import (
	"context"
	"time"

	"github.com/assimon/luuu/config"
	"github.com/assimon/luuu/util/log"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func RedisInit() {
	options := redis.Options{
		Addr:        config.RedisHost + ":" + config.RedisPort,
		DB:          config.RedisDB,
		Password:    config.RedisPassword,
		PoolSize:    config.RedisPoolSize,
		MaxRetries:  config.RedisMaxRetries,
		IdleTimeout: config.RedisIdleTimeout,
	}

	Rdb = redis.NewClient(&options)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Rdb.Ping(ctx).Err(); err != nil {
		log.Sugar.Errorf("[store_redis] Redis connection failed: %v", err)
		panic(err)
	}

	log.Sugar.Info("[store_redis] Redis connection established successfully")
}
