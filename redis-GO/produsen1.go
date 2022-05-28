package main

import (
    "fmt"
    "github.com/go-redis/redis"
	"strconv"
)


func main() {
    fmt.Println("Testing Golang Redis")

    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)

	i := 0
	for {
		name := client.Set("nama", "ramma" + strconv.Itoa(i) , 0).Err()
		if name != nil{
			fmt.Println(name)
		}
		i += 1
	}

	
 }