package redis_utilities

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	pipelineCtx context.Context
)

func setPipelineCtx() {
	pipelineCtx = context.Background()
}

func PipelinedTask(client *redis.Client, dataContext string) (err error) {
	setPipelineCtx()

	if dataContext != "player" {
		return err
	}

	_, err = client.Pipelined(pipelineCtx, pipelinedTaskCallback)

	if err != nil {
		return fmt.Errorf("pipelined failed: %w", err)
	}

	log.Printf("[DOHAN] <info> Player 7's score: %s, challenges completed: %s\n",
		client.HGet(pipelineCtx, "player:7", "score").Val(),
		client.HGet(pipelineCtx, "player:7", "challenges_completed").Val())
	log.Printf("[DOHAN] <info> Player 8's score: %s, challenges completed: %s\n",
		client.HGet(pipelineCtx, "player:8", "score").Val(),
		client.HGet(pipelineCtx, "player:8", "challenges_completed").Val())
	log.Printf("[DOHAN] <info> Player 9's score: %s, challenges completed: %s\n",
		client.HGet(pipelineCtx, "player:9", "score").Val(),
		client.HGet(pipelineCtx, "player:9", "challenges_completed").Val())

	return err
}

func pipelinedTaskCallback(pipe redis.Pipeliner) (err error) {
	err = pipe.HSet(pipelineCtx, "player:7", "score", 15, "challenges_completed", 1).Err()
	if err != nil {
		return err
	}
	err = pipe.HSet(pipelineCtx, "player:8", "score", 18, "challenges_completed", 1).Err()
	if err != nil {
		return err
	}
	err = pipe.HSet(pipelineCtx, "player:9", "score", 12, "challenges_completed", 1).Err()

	return err
}
