package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/robinjoseph08/redisqueue/v2"
	"time"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,  // use default DB
	})

	c, err := redisqueue.NewConsumerWithOptions(&redisqueue.ConsumerOptions{
		VisibilityTimeout: 60 * time.Second,
		BlockingTimeout:   5 * time.Second,
		ReclaimInterval:   1 * time.Second,
		BufferSize:        100,
		Concurrency:       10,
		RedisClient: redisClient,
	})
	if err != nil {
		panic(err)
	}

	c.Register("redisqueue:test", process)

	go func() {
		for err := range c.Errors {
			// handle errors accordingly
			fmt.Printf("err: %+v\n", err)
		}
	}()

	fmt.Println("starting")
	c.Run()
	fmt.Println("stopped")
}

func process(msg *redisqueue.Message) error {
	fmt.Printf("messageid: %v\n", msg.ID)
	fmt.Printf("message: %v\n", msg.Stream)
	fmt.Printf("processing message: %v, %v \n", msg.Values["index"], msg.Values["name"])
	return nil
}
