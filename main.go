package main

import (
	"context"
	"log"
	"time"
	"github.com/go-redis/redis/v8"
	"os"
)

func main() {
	ctx := context.Background()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: ":6379",
	// })

	redisUrl := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Printf("Failed to setup redis, err: %[1]v", err.Error())
		return
	} else {
		opts.PoolSize = 35
		opts.MinIdleConns = 35
		opts.ReadTimeout = 9 * time.Second
		if opts.TLSConfig != nil {
			opts.TLSConfig.InsecureSkipVerify = true
		}
		rdb := redis.NewClient(opts)
		rdb.Set(ctx, "First value", "value_1", 0)
		
		val1, err := rdb.Get(ctx, "value_1").Result()
		if err == redis.Nil {
			log.Println("value_1 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			log.Println("value_1", val2)
		}

		rdb.Set(ctx, "Second value", "value_2", 0)
		val2, err := rdb.Get(ctx, "value_2").Result()
		if err == redis.Nil {
			fmt.Println("value_2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("value_2", val2)
		}
	}
}
