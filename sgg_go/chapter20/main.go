package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, _ := redis.Dial("tcp", "localhost:6379")
	c.Do("Set", "key1", 998)
	// r, _ := redis.Int(c.Do("get", "key1"))
	c.Do("HSet", "user01", "age", 18)
	c.Do("HSet", "user01", "name", "John")
	r2, _ := redis.String(c.Do("HGet", "user01", "name"))
	r3, _ := redis.Int(c.Do("HGet", "user01", "age"))
	fmt.Println(r2, r3)
	c.Close()
	var pool = &redis.Pool{
		MaxIdle:     2,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	c2 := pool.Get()
	defer c2.Close()
	c2.Do("Set", "key1", 998)
	r, _ := redis.Int(c2.Do("get", "key1"))
	fmt.Println(r)
	pool.Close()
}
