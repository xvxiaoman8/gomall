package redis

import (
	"context"
	redislock "github.com/jefferyjob/go-redislock"
	"github.com/redis/go-redis/v9"
	"github.com/xvxiaoman8/gomall/app/checkout/conf"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

// NewRedisLock 创建分布式redis锁
func NewRedisLock(key string, ctx context.Context, opts ...redislock.Option) redislock.RedisLockInter {
	lock := redislock.New(ctx, RedisClient, key)
	if lock == nil {
		panic("redis lock get failed")
	}
	return lock
}

func RedisDo(ctx context.Context, args ...any) (value interface{}, err error) {
	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	// 正确调用 RedisClient.Do 方法
	cmd := RedisClient.Do(ctx, args...)
	value, err = cmd.Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}
