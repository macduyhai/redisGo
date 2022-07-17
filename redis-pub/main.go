package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Publish(ctx, "mychanel-1", "simon message").Err()
	if err != nil {
		log.Fatalf("Pub error-%v\n", err)
	} else {
		log.Println("Pub done")
	}
}
