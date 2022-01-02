package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func postKeys(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	c.String(200, "The key", key, "is successfully saved the value", val)
}

func getValueByKey(c *gin.Context) {
	key := c.Query("key")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	c.String(200, "The value for key", key, "is", val)
}

func main() {
	router := gin.Default()
	router.GET("/getkeys", getValueByKey)
	router.POST("/postkeys", postKeys)
	router.Run("localhost:8080")
}
