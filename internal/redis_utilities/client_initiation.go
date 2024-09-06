package redis_utilities

import "github.com/redis/go-redis/v9"

func NewClient(conn string, db int) (redisClient *redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr:     conn,
		DB:       db,
		Password: "",
	})
}
