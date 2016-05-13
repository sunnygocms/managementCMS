package controllers

import (
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}
type Cmsuser struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Pwd      string `form:"pwd"`
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (this *LoginController) Post() {
	//log := logs.GetLogger()
	//log.SetLogger("file",`{"filename":"post.log"}`)
	//log.EnableFuncCallDepth(true)
	//log.SetLogFuncCallDepth(3)
	beego.Info(this.Input().Get("username"))
	u := Cmsuser{}
	//beego.Info("hhjhjhjh-------"+this.Ctx.Request.ParseForm())
	if err := this.ParseForm(&u); err != nil {
		beego.Info(err)
	} else {
		//this.Ctx.WriteString(u.Username + u.Pwd)
		//log.Debug(err.Error())
	}
	//this.Ctx.WriteString("ssss")
	this.TplName = "index/login.html"
}
