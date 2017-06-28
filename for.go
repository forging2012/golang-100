// FILENAME: 005-for.go
// DATE: 2017/6/11
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: for.go

package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		//i = i + 1
		i++
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}

/**
"C:\Program Files\JetBrains\Gogland 171.4424.55\bin\runnerw.exe" C:/Go\bin\go.exe run C:/Users/admin/Documents/wwwgit/golang-100/for.go
1
2
7
8
9
loop

Process finished with exit code 0
*/