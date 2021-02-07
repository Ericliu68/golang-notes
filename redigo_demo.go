package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.110.118:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Do("set", "test1", "redis")
	v1, err := redis.String(conn.Do("get", "test1"))
	if err != nil {
		fmt.Printf("err is %v\n", err)
	} else {
		fmt.Printf("%s id %s\n", "test1", v1)
	}
}
