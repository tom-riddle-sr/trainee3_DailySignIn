package redis

import (
	"context"
	"trainee3/config"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func New(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Addr,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		logrus.Fatal("連接到redis 失敗:", err)
		return client, err
	}
	logrus.Info("連接到redis 成功")
	return client, nil
}
