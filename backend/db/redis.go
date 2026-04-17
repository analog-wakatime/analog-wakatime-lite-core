package db

import (
	"analog-wakatime-lite-core/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

const redisOpTimeout = 2 * time.Second

func redisCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), redisOpTimeout)
}

func InitRedis() error {
	addr := config.GetRedisURL()
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx, cancel := redisCtx()
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func SetDeviceSession(key string, data string, expiration time.Duration) error {
	ctx, cancel := redisCtx()
	defer cancel()

	return RDB.Set(ctx, "device_session:"+key, data, expiration).Err()
}
func GetDeviceSession(key string) (string, error) {
	ctx, cancel := redisCtx()
	defer cancel()

	return RDB.Get(ctx, "device_session:"+key).Result()
}

func DeleteDeviceSession(key string) error {
	ctx, cancel := redisCtx()
	defer cancel()

	return RDB.Del(ctx, "device_session:"+key).Err()
}
