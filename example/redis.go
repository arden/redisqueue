package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/util/gconv"
)

type User struct {
	ID           uint   `json:"id"`            // 主键
	Name 		 string `json:"user_name"` 				   // 转为json的字符串
}

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	user := &User{
		ID: 1000,
		Name: "arden",
	}

	userMap := gconv.Map(user)

	result := redisClient.HMSet(context.Background(), "test_user", userMap)
	println(result.Val())

	resultUser := redisClient.HGetAll(context.Background(), "test_user")
	println(resultUser.String())
	var destUser *User
	_ = gconv.Struct(resultUser.Val(), &destUser)
	println(destUser.Name)

	redisClient.ZAdd(context.Background(), "realtimeInvitedUserRedisKey", &redis.Z{
		Score: gconv.Float64(1000),
		Member: gconv.String(2000),
	})

	redisClient.ZIncrBy(context.Background(), "aaa", 1, gconv.String(1000))

}
