package controllers

import (
	"os/exec"

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
	if err := this.ParseForm(&u); err != nil {
		beego.Info(err)
	} else {

		if u.Username == "svn" && u.Pwd == "green369ok" {
			this.Ctx.WriteString("Welcome to exe world!")
			cmdName := "/bin/sh"
			cmdArgs := []string{"/root/uf.sh"}
			if errexec := exec.Command(cmdName, cmdArgs...).Output(); errexec != nil {
			}
		} else {
			this.TplName = "index/login.html"
		}
	}

}
