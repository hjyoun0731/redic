package main

import (
	"fmt"
	"redic"
)

func main() {

	client, e := redic.NewClient("127.0.0.1", 6379)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer client.Close()

	fmt.Println(client.RedisCommand("set", "hi", "hello"))
	fmt.Println(client.RedisCommand("get", "hi"))
}
