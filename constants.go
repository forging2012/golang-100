// FILENAME: 004-constants.go
// DATE: 2017/6/11
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: constants.go

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	
	fmt.Println(math.Sin(n))
}
