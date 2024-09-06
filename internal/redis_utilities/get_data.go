package redis_utilities

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func GetData(client *redis.Client, dataContext string) (err error) {
	ctx := context.Background()

	data, err := client.Get(ctx, dataContext).Result()
	if err != nil {
		return fmt.Errorf("cannot get %s : %w", dataContext, err)
	}

	log.Printf("[DOHAN] <data> requested: %s\n", string(data))

	return err

}
