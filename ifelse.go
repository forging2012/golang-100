// FILENAME: 006-ifelse.go
// DATE: 2017/6/11
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: ifelse.go
package main

import (
	"fmt"
)

func main() {
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4.")
	}

	if 7%2 == 0 {
		fmt.Println("7 is even.")
	} else {
		fmt.Println("7 is odd.")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	
}

/**
"C:\Program Files\JetBrains\Gogland 171.4424.55\bin\runnerw.exe" C:\Go\bin\go.exe run C:/Users/admin/Documents/wwwgit/golang-100/ifelse.go
8 is divisible by 4.
7 is odd.
9 has 1 digit

Process finished with exit code 0

*/