package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Endpoint string `yaml:"endpoint"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewClient(cfg RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Endpoint,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return client
}

func CloseClient(client *redis.Client) {
	defer client.Close()
}
