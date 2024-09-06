package internal

import (
	"github.com/redis/go-redis/v9"
	"golang-redis/internal/redis_utilities"
	"log"
	"os"
)

func ShowUsage() {
	log.Println(`[DOHAN] <warning> go run main.go <demo>
Where <demo> is one of:
	- ping ==> run the ping command
	- get <dataContext> ==> get a string value
	- expired-key <dataContext> ==> set an expiring key
	- pipeline <dataContext> ==> run a batch of commands
	- transaction ==> run a batch of commands that must all succeed
	- pubsub ==> send a message to a channel and listen to the channel
	- reset ==> restore the initial data set`)
}

func Boot(dbConn string, db int) (err error) {
	client := redis_utilities.NewClient(dbConn, db)

	err = switchUtilities(client)

	return err
}

func switchUtilities(client *redis.Client) (err error) {
	log.Println("[DOHAN] <info> Processing...")

	switch os.Args[1] {
	case "ping":
		err = redis_utilities.PingClient(client)
		if err == nil {
			log.Println("[DOHAN] <exit> OK!")
		}
	case "get":
		err = redis_utilities.GetData(client, string(os.Args[2]))
		if err == nil {
			log.Println("[DOHAN] <exit> OK!")
		}

	case "expired-key":
		err = redis_utilities.ExpiredKeyTime(client, string(os.Args[2]))
		if err == nil {
			log.Println("[DOHAN] <exit> OK!")
		}
	case "pipeline":
		err = redis_utilities.TaskPipelined(client, string(os.Args[2]))
		if err == nil {
			log.Println("[DOHAN] <exit> OK!")
		}
	default:
		log.Println("[DOHAN] <exit> Default case!")
		ShowUsage()
	}

	return err
}
