package controllers

type IndexController struct {
	BaseController
}

// @router /index
func (this *IndexController) Index() {
	this.Ctx.WriteString(this.GetSession("username").(string))
	this.Html("aaaaaaa")
}

func (this *IndexController) Get() {
	this.Html("aaaaaaa")
}
