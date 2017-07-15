package main

import (
	//_ "quickstart/routers"
	//"quickstart/controllers"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "kexuekong@qq.com"
	this.TplName = "index.tpl"
}
func main() {
	//beego.Run()
	beego.Router("/",&controllers.MainController{})
	beego.Router("/main",&controllers.MainController{})
	beego.Router("/home",&controllers.MainController{})
	//beego.Router("/user",&controllers.UserController{})
}