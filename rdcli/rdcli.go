package rdcli

import (
	"context"
	"errors"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/macduyhai/redisGo/models"
)

type RedisCli struct {
	Client *redis.Client
	// Client *redis.ClusterClient
}

var (
	ctx          = context.Background()
	rediscli     *RedisCli
	one          sync.Once
	sentinelHost = "192.168.2.9"
)

// Single ton pattern
func NewRedisClient(add string) (*RedisCli, error) {
	one.Do(func() {
		rediscli = &RedisCli{}
		rediscli.Client = redis.NewFailoverClient(&redis.FailoverOptions{
			// rediscli.Client = redis.NewClusterClient(&redis.ClusterOptions{
			MasterName:    "mymaster",
			SentinelAddrs: []string{sentinelHost + ":26379", sentinelHost + ":26380", sentinelHost + ":26381"},
			Password:      "admin",
			DB:            0,
		})

		// Client init basic
		// rediscli.Client = redis.NewClient(&redis.Options{
		// 	Addr:     add,
		// 	Password: "admin",
		// 	DB:       0,
		// })
	})
	if err := rediscli.Client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rediscli, nil

}

func (rediscli *RedisCli) SaveRedisDB(data interface{}) error {
	switch (reflect.TypeOf(data)).String() {
	case "models.User":
		user := data.(models.User)
		if user.UserName == "" {
			return errors.New("username not emty")
		}
		if user.Password == "" {
			return errors.New("password not emty")
		}
		if user.FullName == "" {
			return errors.New("fullname not emty")
		}
		jsonuser, err := json.Marshal(user)
		if err != nil {
			return err
		}
		// Check user
		_, err = rediscli.Client.Get(ctx, user.UserName).Result()
		if err == nil {
			return errors.New("user exited")
		} else if (reflect.TypeOf(err)).String() == "*net.OpError" {
			return errors.New("disconected to server redis")
		}
		// fmt.Printf("Type: %T, err: %v\n", err, err)
		rediscli.Client.Set(ctx, user.UserName, jsonuser, 24*time.Hour) // time : thoi gian ton tai
	default:
		log.Println("Used default")
	}
	return nil
}
func (rediscli *RedisCli) LoadAllData() {
	keys, err := rediscli.Client.Do(ctx, "KEYS", "*").Result()
	if err != nil {
		log.Println("Load data Error")
	} else {
		log.Printf("List keys DB: %v", keys)
	}

}
