package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	redisConn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	set := redisConn.Set(ctx, "name", "King", 10*time.Minute)
	fmt.Println(set)
	get := redisConn.Get(ctx, "name")
	fmt.Println(get)
}
