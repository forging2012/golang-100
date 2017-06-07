// FILENAME: 101-beego-hello-world.go
// DATE: 2017/6/8
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: 101-beego-hello-world.gopackage main

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