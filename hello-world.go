// FILENAME: 001-hello-world.go
// DATE: 2017/6/8
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: hello-world.go

package main
import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
}


//GOROOT=C:/Go
//GOPATH=C:/Go/bin
//C:/Go\bin\go.exe build -i -o C:\Users\admin\AppData\Local\Temp\Build_001_hello_world_go_and_rungo -gcflags "-N -l" C:/Users/admin/Documents/wwwgit/golang-100/hello-world.go
//"C:\Program Files\JetBrains\Gogland 171.4424.55\bin\runnerw.exe" "C:\Program Files\JetBrains\Gogland 171.4424.55\plugins\intellij-go-plugin\lib\dlv\windows\dlv.exe" --listen=localhost:56296 --headless=true --api-version=2 --backend=default exec C:\Users\admin\AppData\Local\Temp\Build_001_hello_world_go_and_rungo --
//API server listening at: 127.0.0.1:56296
//Hello world!
//
//Process finished with exit code 0
