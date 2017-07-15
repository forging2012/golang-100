// FILENAME: dht-getpeers-redis-v4.go
// DATE: 2017/7/9
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: dht-getpeers-redis-v4.go

package main

import (
	"fmt"
	"github.com/shiyanhui/dht"
	"time"
	"gopkg.in/redis.v4"
)

func main() {


	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set  DB: 0, // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err) // Output: PONG <nil>

	err = client.Set("name", "xys", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)



	d := dht.New(nil)
	go d.Run()

	for {
		// ubuntu-14.04.2-desktop-amd64.iso
		peers, err := d.GetPeers("546cf15f724d19c4319cc17b179d7e035f89c1f4")
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		//err = client.Set("peers", peers, 100).Err()
		//err = client.Set("peers", peers, 0).Err()
		//err = client.RPush("peers", peers).Err()
		//client.HSet("peers", "peers", "peers");
		//err = client.RPush(fmt.Sprintf("%s",peers), fmt.Sprintf("%s",peers), 0).Err()
		//err = client.RPush("peers", fmt.Sprintf("%s",peers), 0).Err()

		if err != nil {
			panic(err)
		}
		//fmt.Println("Found peers:", peers)
		fmt.Println(peers)
		//fmt.Printf("%s", peers)
		fmt.Sprintf("%s",peers)

	}
}
