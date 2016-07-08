package routers

import (
	"github.com/astaxie/beego"
	"github.com/sunnygocms/managementCMS/controllers"
)

func init() {
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.EditorController{})
	beego.AutoRouter(&controllers.UsergroupController{})
	beego.AutoRouter(&controllers.PowerController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.Router("/svn", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	//beego.Router("/index/:action([A-za-z]+)/?:id([0-9]+)/", &controllers.IndexController{}, "*:Index")

}
