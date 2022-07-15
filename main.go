package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Deverloper struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//-------------------------------

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	// Data type struct
	json, err := json.Marshal(Deverloper{Name: "Mac Duy Hai", Age: 27})
	if err != nil {
		fmt.Println(err)
	}

	err = rdb.Set(ctx, "id1234", json, 3*time.Second).Err()
	if err != nil {
		panic(err)
	}
	vals, err := rdb.Get(ctx, "id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vals)
	time.Sleep(4 * time.Second)
	vals, err = rdb.Get(ctx, "id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vals)

	// Data type key, value
	err = rdb.Set(ctx, "name", "Simon", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Value:", val)
	// Data not exits
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

}
func main() {
	ExampleClient()
}
