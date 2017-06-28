// FILENAME: hello-world-with-route.go
// DATE: 28/06/2017
// AUTHOR: FORGING2012@GMAIL.COM
// Github: https://github.com/forging2012
// Description: hello-world-with-route.go

package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello world.")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}