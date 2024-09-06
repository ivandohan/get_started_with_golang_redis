package redis_utilities

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ExpiredKeyTime(client *redis.Client, dataContext string) (err error) {
	ctx := context.Background()

	err = client.HSet(
		ctx,
		dataContext,
		"name", "Crymyios",
		"score", 0,
		"team", "Knucklewimp",
		"challenges_completed", 0).Err()

	if err != nil {
		return fmt.Errorf("cannot set new %s: %w", dataContext, err)
	}

	if !client.Expire(ctx, dataContext, time.Second).Val() {
		return fmt.Errorf("cannot set expired time for %s : %w", dataContext, err)
	}

	for i := 0; i < 3; i++ {
		val, err := client.HGet(ctx, dataContext, "name").Result()
		if err != nil {
			log.Printf("[DOHAN] <info> %s has expired.\n", dataContext)
			return nil
		}

		log.Printf("[DOHAN] <data> %s's name : %s\n", dataContext, val)
		time.Sleep(500 * time.Millisecond)
	}

	return err
}
