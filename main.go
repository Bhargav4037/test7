package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type cache struct {
	first string
}

func (c cache) InitPool(redisHost, redisPort string) {
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	msg := fmt.Sprintf("Initialized Redis at %s", redisAddr)
	const maxConnections = 10
	redis.DialDatabase(2)
	log.Println(msg)
}

func main() {
	c := cache{"heyy.."}
	c.InitPool("abcd", "1234")
}
