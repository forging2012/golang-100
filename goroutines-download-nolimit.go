package main

import (
	"fmt"
	"strconv"
	"time"
)

// 模拟下载页面的方法
func download(url string) {
	fmt.Println("download from ", url)
}

func main() {
	urls := [100]string{}
	for i := 0; i < 100; i++ {
		urls[i] = "url" + strconv.Itoa(i)
	}
	for i := 0; i < len(urls); i++ {
		go download(urls[i])
	}

	// 休眠一下
	for {
		time.Sleep(1 * 1e9)
	}
}