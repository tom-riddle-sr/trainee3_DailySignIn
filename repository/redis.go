package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	Set(c *redis.Client, key string, value interface{}, expirationTime int32) error
	Get(c *redis.Client, key string, model interface{}) error
}

type Redis struct{}

func NewRedis() IRedis {
	return &Redis{}
}

func (r *Redis) Set(c *redis.Client, key string, value interface{}, expirationTime int32) error {
	ctx := context.Background()
	err := c.Set(ctx, key, value, time.Duration(expirationTime)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(c *redis.Client, key string, model interface{}) error {
	ctx := context.Background()
	val, err := c.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(val), model); err != nil {
		return err
	}
	return nil
}
