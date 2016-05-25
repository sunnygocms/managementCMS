package controllers

//"github.com/astaxie/beego"

type IndexController struct {
	BaseController
}

// @router /index
func (self *IndexController) Index() {
	sess_username, _ := self.GetSession("editor_username").(string)

	if len(sess_username) == 0 {
		self.Html(sess_username + "00000000000-------")
		self.Ctx.Redirect(302, "/login")
	} else {
		sess_power, _ := self.GetSession("editor_power").(string)
		self.Data["editor_username"] = sess_username
		self.Data["editor_power"] = sess_power
		self.TplName = "index/index.html"
	}

}
func (self *IndexController) Welcome() {
	self.TplName = "index/welcome.html"
}
