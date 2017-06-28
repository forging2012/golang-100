// FILENAME: 102-beego-hello-world-with-route.go
// DATE: 2017/6/8
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: beego-hello-world-with-route.go
package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}