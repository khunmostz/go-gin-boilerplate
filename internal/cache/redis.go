package cache

import (
	"fmt"
	"go-gin-boilerplate/config"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis(config *config.RedisConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		panic(err)
	}

	log.Println("Successfully connected to Redis!")

	return redisClient
}
