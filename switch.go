// FILENAME: 007-switch.go
// DATE: 2017/6/11
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: switch.go
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	case t.Hour() >= 12:
		fmt.Println("It's after noon.")
	}

}
