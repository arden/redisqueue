package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/robinjoseph08/redisqueue/v2"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,  // use default DB
	})

	p, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      10000,
		ApproximateMaxLength: true,
		RedisClient: redisClient,
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		err := p.Enqueue(&redisqueue.Message{
			Stream: "redisqueue:test",
			Values: map[string]interface{}{
				"index": i,
				"name": fmt.Sprintf("arden:%v", i),
			},
		})
		if err != nil {
			panic(err)
		}

		if i%100 == 0 {
			fmt.Printf("enqueued %d\n", i)
		}
	}
}
