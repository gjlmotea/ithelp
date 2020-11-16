package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	c := NewClient()
	test(c)
}
func NewClient() *redis.Client { // 實體化redis.Client 並返回實體的位址
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

func test(c *redis.Client) { // 對該 redis.Client 進行操作
	err := c.Set("key", "value", 0).Err() // => SET key value 0 數字代表過期秒數，在這裡0為永不過期
	if err != nil {
		panic(err)
	}

	val, err := c.Get("key").Result() // => GET key
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := c.Get("key2").Result() // => GET key2
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
