// FILENAME: 003-variables.go
// DATE: 2017/6/11
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: variables.go

package main

import (
	"fmt"
)

func main() {
	var a string = "This is a test text."
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	
	f := "short"
	fmt.Println(f)
	
}
