package controllers

import (
	"os/exec"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	//	"github.com/astaxie/beego/logs"
)

var cpt *captcha.Captcha

type LoginController struct {
	beego.Controller
}
type Cmsuser struct {
	Id         int    `form:"-"`
	Username   string `form:"username"`
	Pwd        string `form:"pwd"`
	Verifycode string `form:"captcha"`
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}
func (this *LoginController) Index() {

	this.TplName = "index/login.html"

}
func (c *LoginController) Get() {

	c.TplName = "index.tpl"
}
func (this *LoginController) Post() {
	//log := logs.GetLogger()
	//log.SetLogger("file",`{"filename":"post.log"}`)
	//log.EnableFuncCallDepth(true)
	//log.SetLogFuncCallDepth(3)
	//	beego.Info(this.Input().Get("username"))
	u := Cmsuser{}
	if err := this.ParseForm(&u); err != nil {
		beego.Info(err)
	} else {

		if u.Username == "svn" && u.Pwd == "green369ok" {
			if errexec := exec.Command("/bin/sh", "/root/uf.sh").Run(); errexec != nil {

			} else {
				this.Ctx.WriteString("Welcome to exe world!")
			}
		} else {
			//			id, value := this.GetString("captcha_id"), u.Verifycode

			//			if len(value) == 0 {
			//				this.Index()
			//			} else {
			//				this.Ctx.WriteString("v:" + value + "   captcha:" + id)
			//			}
			if !cpt.VerifyReq(this.Ctx.Request) {
				this.Index()
			}

		}
	}

}
