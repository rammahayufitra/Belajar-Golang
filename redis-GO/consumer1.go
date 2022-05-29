package main
import (
	"fmt"
	"github.com/go-redis/redis"
	)

func main(){
	fmt.Println("Testing Golang Redis-Consumer")

    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)

	i := 0
	for {
		name, err:= client.Get("nama").Result()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(name)
		i += 1
	}
}