package main

import (
	"golang-redis/internal"
	"log"
	"os"
)

const (
	db     = 0
	dbConn = "localhost:6379"
)

func main() {
	if len(os.Args) < 1 {
		internal.ShowUsage()
		os.Exit(0)
	}

	err := internal.Boot(dbConn, db)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
