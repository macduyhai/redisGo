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
	sub := rdb.Subscribe(ctx, "mychanel-1")
	defer sub.Close()
	// for {
	// 	msg, err := sub.ReceiveMessage(ctx)
	// 	if err != nil {
	// 		log.Fatalf("receive error-%v", err)
	// 	}
	// 	log.Println(msg.Channel, msg.Payload)
	// }
	// Su dung chanel
	ch := sub.Channel()
	for msg := range ch {
		log.Println(msg.Channel, msg.Payload)
	}
}
