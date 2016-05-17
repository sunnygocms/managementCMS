package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "github.com/sunnygocms/managementCMS"
	c.Data["Email"] = "jinheking@gmail.com"
	c.TplName = "index.tpl"
}
