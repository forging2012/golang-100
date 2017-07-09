// FILENAME: redis-demo.go
// DATE: 2017/7/9
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: golang-redis-v4.go

package main
import (
	"fmt"
	"gopkg.in/redis.v4"
)



func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set  DB: 0, // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err) // Output: PONG <nil>

	//fmt.Println("Golang redis.")
}