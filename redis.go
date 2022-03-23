package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func testRedis() {
	log.Println("Testing REDIS...")

	redisUrl := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASS")

	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPass,
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = client.Set(ctx, "name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	saved, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Returned:", saved)
	log.Println("Test finished!")
}
