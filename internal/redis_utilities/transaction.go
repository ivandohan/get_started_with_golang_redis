package redis_utilities

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	transactionCtx context.Context
)

func setTransactionCtx() {
	transactionCtx = context.Background()
}

func Transaction(client *redis.Client) (err error) {

	setTransactionCtx()

	_, err = client.TxPipelined(transactionCtx, pipelinedTaskCallback)
	if err != nil {
		return fmt.Errorf("TxPipelined failed: %w", err)
	}

	log.Printf("[DOHAN] <info> Sykios's new team: %s\n", client.HGet(transactionCtx, "player:1", "team").Val())
	log.Printf("[DOHAN] <info> Belaeos's new team: %s\n", client.HGet(transactionCtx, "player:4", "team").Val())
	log.Printf("[DOHAN] <info> Tiaitia's new team: %s\n", client.HGet(transactionCtx, "player:3", "team").Val())
	log.Printf("[DOHAN] <info> Team Grumblebum: %s\n", client.SMembers(transactionCtx, "team:Grumblebum").Val())
	log.Printf("[DOHAN] <info> Team Knucklewimp: %s\n", client.SMembers(transactionCtx, "team:Knucklewimp").Val())

	return err
}

func transactionPipelinedCallback(pipe redis.Pipeliner) (err error) {
	// Move Sykios to team Grumblebum
	err = pipe.HSet(transactionCtx, "player:1", "team", "Grumblebum").Err()
	if err != nil {
		return err
	}
	// Move Nidios to team Grumblebum
	err = pipe.HSet(transactionCtx, "player:2", "team", "Grumblebum").Err()
	if err != nil {
		return err
	}
	// Move Belaeos to team Grumblebum
	err = pipe.HSet(transactionCtx, "player:4", "team", "Grumblebum").Err()
	if err != nil {
		return err
	}
	// Move Tiaitia to team Knucklewimp
	err = pipe.HSet(transactionCtx, "player:3", "team", "Knucklewimp").Err()
	if err != nil {
		return err
	}

	// Team update: remove Belaeos from team Knucklewimp
	err = pipe.SRem(transactionCtx, "team:Knucklewimp", "Belaeos").Err()
	if err != nil {
		return err
	}

	// Team update: add Tiaitia to team Knucklewimp
	err = pipe.SAdd(transactionCtx, "team:Knucklewimp", "Tiaitia").Err()
	if err != nil {
		return err
	}

	// Add team Grumblebum
	err = pipe.SAdd(transactionCtx, "team:Grumblebum", "Sykios", "Nidios", "Belaeos").Err()
	if err != nil {
		return err
	}

	// Remove team Dorkfoot. A set is removed by removing all elements.
	err = pipe.SRem(transactionCtx, "team:Dorkfoot", "Sykios", "Nidios", "Tiaitia").Err()
	if err != nil {
		return err
	}

	return err
}
