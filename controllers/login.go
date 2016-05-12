package controllers

import (
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}
type user struct {
	username string
	pwd	string
}
func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (this *LoginController) Post(){
	//log := logs.GetLogger()	
	//log.SetLogger("file",`{"filename":"post.log"}`)
	//log.EnableFuncCallDepth(true)
	//log.SetLogFuncCallDepth(3)
	u := user{}
	this.Ctx.Request.ParseForm()
	if err := this.ParseForm(&u);err !=nil{
		this.Ctx.WriteString(u.username+u.pwd)
	}else{
		beego.Info(err)
		//log.Debug(err.Error())
		
	}
	this.Ctx.WriteString("ssss")
}
