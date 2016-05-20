package routers

import (
	"github.com/astaxie/beego"
	"github.com/sunnygocms/managementCMS/controllers"
)

func init() {
	beego.AutoRouter(&controllers.IndexController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	//beego.Router("/index/:action([A-za-z]+)/?:id([0-9]+)/", &controllers.IndexController{}, "*:Index")

}
