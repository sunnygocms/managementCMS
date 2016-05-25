package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sunnygocms/managementCMS/models"
)

type BaseController struct {
	beego.Controller
	ControllerName string
	ActionName     string
	TplNames       string
}

func (self *BaseController) Prepare() {
	controller, action := self.GetControllerAndAction()
	sess_username, _ := self.GetSession("editor_username").(string)
	self.Data["viewpath"] = "/static"
	if controller != "LoginController" && action != "Get" {
		if sess_username == "" {
			self.Ctx.Redirect(302, "/login")
		}

	}
}

// @Title Get
// @Description get SunnyEditor by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunnyEditor
// @Failure 403 :id is empty
// @router /:id [get]
func (this *BaseController) GetOne() {
	//	idStr := this.Ctx.Input.Param(":id")
	//	id, _ := strconv.Atoi(idStr)
	//	v, err := models.GetSunnyEditorById(id)
	//	if err != nil {
	//		this.Data["json"] = err.Error()
	//	} else {
	//		this.Data["json"] = v
	//	}
	//	this.ServeJSON()
}

//直接在页面输出字符串
func (this *BaseController) Html(str string) {
	this.Ctx.WriteString(str) //self.GetControllerAndAction()
}
