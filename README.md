This repo is a fork of [https://github.com/redis/hiredis][1].

---

![gopher](https://user-images.githubusercontent.com/55075487/146725496-5c07d4bc-bac9-4856-97a4-19afb3b632b7.jpg)

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
