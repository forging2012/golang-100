// FILENAME: dht-getpeers.go
// DATE: 2017/7/9
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: dht-getpeers.go

package main

import (
	"fmt"
	"github.com/shiyanhui/dht"
	"time"
)

func main() {
	d := dht.New(nil)
	go d.Run()

	for {
		// ubuntu-14.04.2-desktop-amd64.iso
		peers, err := d.GetPeers("546cf15f724d19c4319cc17b179d7e035f89c1f4")
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		fmt.Println("Found peers:", peers)
	}
}
