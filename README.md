This repo is a fork of [https://github.com/redis/hiredis][1].

---

![20211220_155815](https://user-images.githubusercontent.com/55075487/146725244-3a5cc056-1115-438b-8fcc-96de44bdc465.jpg)

redic - Bindings for hiredis Redis-client Go library

## Install

```sh
go get -u github.com/hjyoun0731/redic
```

## Example

```
package main

import (
	"fmt"

	"github.com/hjyoun0731/redic"
)

func main() {

	client, e := redic.NewClient("127.0.0.1", 6379)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer client.Close()

	fmt.Println(client.RedisCommand("set", "hi", "hello")) // print : OK
	fmt.Println(client.RedisCommand("get", "hi"))          // print : hello
}
```

[hjyoun0731/redic/example][2].


[1]: https://github.com/redis/hiredis
[2]: https://github.com/hjyoun0731/redic/blob/master/example/main.go
