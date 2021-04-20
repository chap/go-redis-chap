package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	stop := runExporter(ctx)
	defer stop(ctx)

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
		rdb = redis.NewClient(opts)
	}


	rdb.Set(ctx, "First value", "value_1", 0)

	rdb.Set(ctx, "Second value", "value_2", 0)
}
