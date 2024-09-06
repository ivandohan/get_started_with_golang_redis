package redis_utilities

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func PingClient(client *redis.Client) (err error) {
	ctx := context.Background()

	log.Println("[DOHAN]", client.Ping(ctx))

	info, err := client.ClientInfo(ctx).Result()
	if err != nil {
		return fmt.Errorf("method ClientInfo failed: %w", err)
	}

	log.Printf("[DOHAN] <info> %#v\n", info)

	return err
}
