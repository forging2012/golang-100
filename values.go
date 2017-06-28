// FILENAME: 002-values.go
// DATE: 2017/6/11
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: values.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+2" + "=",1+2 )
	fmt.Println("7/2" + "=",7/2 )
	fmt.Println("7.0/2" + "=",7.0/2 )
	fmt.Println("7.0/2.0" + "=",7.0/2.0 )
	fmt.Println("7/2.0" + "=",7/2.0 )
	fmt.Println("7.0/3.5" + "=",7.0/3.5 )
	
	fmt.Println( true && true )
	fmt.Println( true || true )
	fmt.Println( true && false )
	fmt.Println( true || false )
	fmt.Println( false || false )
	fmt.Println( false && false )
	fmt.Println( ! true )
	fmt.Println( ! false )

}


/*
"C:\Program Files\JetBrains\Gogland 171.4424.55\bin\runnerw.exe" C:/Go\bin\go.exe run C:/Users/admin/Documents/wwwgit/golang-100/values.go
golang
1+2= 3
7/2= 3
7.0/2= 3.5
7.0/2.0= 3.5
7/2.0= 3.5
7.0/3.5= 2
true
true
false
true
false
false
false
true

Process finished with exit code 0

*/
