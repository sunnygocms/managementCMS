package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	//	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sunnygocms/managementCMS/models"
)

type BaseController struct {
	beego.Controller
	ControllerName string
	ActionName     string
	TplNames       string
}

func (this *BaseController) Prepare() {
	controller, action := this.GetControllerAndAction()
	sess_username, _ := this.GetSession("editor_username").(string)
	this.Data["viewpath"] = "/static"
	this.Data["Id"] = this.GetEditorId()
	if controller != "LoginController" && action != "Get" {
		if len(sess_username) == 0 {
			this.Ctx.Redirect(302, "/login")
		}
	}
}

//直接在页面输出字符串
func (this *BaseController) Html(str string) {
	this.Ctx.WriteString(str) //self.GetControllerAndAction()
}

//show message
func (this *BaseController) Info(v ...interface{}) {
	beego.Info("========>", v)
}

//get editor id
//It is use to selcet from table;
func (this *BaseController) GetEditorId() int {
	editorid, _ := this.GetSession("editor_userID").(int)
	return editorid
}

//IsPost is this a post method request?
func (this *BaseController) IsPost() bool {
	return this.Ctx.Input.IsPost()
}

//IsSubmit is this is a submit?
func (this *BaseController) IsSubmit() bool {
	keymap := this.Input()
	_, ok := keymap["submit"]
	return ok
}

/*
* Success
 */
func (this *BaseController) Success(msg string, url string, wait int) {
	data := make(map[string]interface{})
	data["type"] = true
	data["title"] = "提示信息"
	data["msg"] = msg
	data["wait"] = wait
	if url == "-1" {
		url = this.Ctx.Request.Referer()
	} else if url == "-2" {
		url = this.Ctx.Request.Referer()
	}
	data["url"] = url
	this.Data["mess"] = data
	this.TplName = "util/message.html"

}

/*
* 失败跳转
 */
func (this *BaseController) Error(msg string, url string, wait int) {
	data := make(map[string]interface{})
	data["type"] = false
	data["title"] = "错误提示"
	data["msg"] = msg
	data["wait"] = wait
	if url == "-1" {
		url = this.Ctx.Request.Referer()
	} else if url == "-2" {
		url = this.Ctx.Request.Referer()
	}

	data["url"] = url
	this.Data["mess"] = data
	this.TplName = "util/message.html"

}

//Ajax return
func (this *BaseController) AjaxReturn(status int, msg string, data interface{}) {
	json := make(map[string]interface{})
	json["status"] = status
	json["msg"] = msg
	json["data"] = data
	this.Data["json"] = json
	this.ServeJSON()
	return
}

//md5 return
func (this *BaseController) SunnyMd5(str string) (keyMd5 string) {
	m := md5.Sum([]byte(str))
	keyMd5 = hex.EncodeToString(m[:])
	return
}

//权限判断
func (this *BaseController) CheckPower(controller string, action string) (result bool) {
	power, err := models.GetEditorPowersById(this.GetEditorId())
	if err != nil {
		result = false
	} else {
		arr := power.(map[string][]string)
		value, isExist := arr[controller]
		result = false
		if isExist {
			for _, a := range value {
				if a == action {
					result = true
					break
				}
			}
		}
	}
	return
}

//func (this *BaseController) GetActionName() {
//	_, action_name := this.GetControllerAndAction()
//	this.Data["ACTION_NAME"] = action_name
//}

func (c *BaseController) SunnyJSON(encoding ...bool) string {
	//	var (
	//		hasIndent   = true
	//		hasEncoding = false
	//	)
	//	if BConfig.RunMode == PROD {
	//		hasIndent = false
	//	}
	//	if len(encoding) > 0 && encoding[0] == true {
	//		hasEncoding = true
	//	}

	//	if hasIndent {
	//		content, err = json.MarshalIndent(data, "", "  ")
	//	} else {
	//		content, err = json.Marshal(data)
	//	}
	//	if err != nil {
	//		http.Error(output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
	//		return err
	//	}
	content, _ := json.MarshalIndent(c.Data["json"], "", "  ")
	s := string(content)
	return s

}
func stringsToJSON(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}
