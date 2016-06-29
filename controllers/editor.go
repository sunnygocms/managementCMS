package controllers

import (
	"encoding/json"
	//	"errors"
	//	"fmt"
	"strconv"
	//	"strings"

	"github.com/sunnygocms/managementCMS/models"
)

// oprations for SunnyEditor
type EditorController struct {
	BaseController
}

type EditorInput struct {
	Id          int    `form:"id"`
	Username    string `form:"username"`
	Password    string `form:"password"`
	Repwd       string `form:"re_password"`
	Usergroup   []int  `form:"usergroup"`
	Description string `form:"description"`
}

func (this *EditorController) URLMapping() {
	this.Mapping("Post", this.Post)
	//	this.Mapping("GetOne", this.GetOne)
	//	this.Mapping("GetAll", this.GetAll)
	//	this.Mapping("Put", this.Put)
	this.Mapping("List", this.List)
	this.Mapping("Edit", this.Edit)
	this.Mapping("Add", this.Add)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create SunnyEditor
// @Param	body		body 	models.SunnyEditor	true		"body for SunnyEditor content"
// @Success 201 {int} models.SunnyEditor
// @Failure 403 body is empty
// @router / [post]
func (this *EditorController) Post() {
	var v models.SunnyEditor
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSunnyEditor(&v); err == nil {
			this.Ctx.Output.SetStatus(201)
			this.Data["json"] = v
		} else {
			this.Data["json"] = err.Error()
		}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

/**
// @Title Get
// @Description get SunnyEditor by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunnyEditor
// @Failure 403 :id is empty
// @router /:id [get]
func (this *EditorController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunnyEditorById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJSON()
}

// @Title Get All
// @Description get SunnyEditor
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunnyEditor
// @Failure 403
// @router / [get]
func (this *EditorController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSunnyEditor(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJSON()
}

// @Title Update
// @Description update the SunnyEditor
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunnyEditor	true		"body for SunnyEditor content"
// @Success 200 {object} models.SunnyEditor
// @Failure 403 :id is not int
// @router /:id [put]
func (this *EditorController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunnyEditor{Id: id}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunnyEditorById(&v); err == nil {
			this.Data["json"] = "OK"
		} else {
			this.Data["json"] = err.Error()
		}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
*/
// @Title Delete
// @Description delete the SunnyEditor
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *EditorController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunnyEditor(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

//get editor list
func (this *EditorController) List() {

	var fields = []string{"Id", "Username", "Description"}
	var limit int64 = 100
	var offset int64 = 0
	l, err := models.GetAllEditor(fields, offset, limit)
	if err != nil {
		this.Data["Editor"] = err.Error()
	} else {
		this.Data["Editor"] = l
	}
	this.Data["Id"] = this.GetEditorId()
	this.TplName = "editor/list.html"
}

//编辑
func (this *EditorController) Edit() {
	if this.IsSubmit() {

	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		usergroup, _ := models.GetAllUserGroup()
		this.Data["Usergroup"] = usergroup
		var id int
		mapp := this.Ctx.Input.Params()
		id, _ = strconv.Atoi(mapp["1"])
		editor, err := models.GetSunnyEditorById(id)
		if err == nil {
			this.Data["Editor"] = &editor
		} else {
			this.Info(err)
		}
		userandgroup, e := models.GetSunnyUserAndGroupById(id)
		if e == nil {
			this.Data["Userandgroup"] = userandgroup
		} else {
			this.Info(err)
		}
		this.TplName = "editor/form.html"
	}
}

//添加
func (this *EditorController) Add() {
	if this.IsSubmit() {
		ei := EditorInput{}
		if err := this.ParseForm(&ei); err != nil {
			this.Info(err)
		} else {
			if len(ei.Repwd) == 0 || len(ei.Password) == 0 {
				this.Error("密码不能够为空", "-1", 4)
			} else if ei.Password != ei.Repwd {
				this.Error("新的密码两次输入不同", "-1", 4)
			}
			//TODO 此处开始完成保存，保存分两部分一部分是 sunny_editor表，基本完成
			//一部分是sunny_user_and_group表
			//删除缓存
			var sunnyeditor models.SunnyEditor
			sunnyeditor.Username = ei.Username
			sunnyeditor.Password = this.SunnyMd5(ei.Password)
			sunnyeditor.Description = ei.Description
			sunnyeditor.Avatar = ""
			sunnyeditor.Status = 1
			id, err := models.AddSunnyEditor(&sunnyeditor)
			if err != nil {
				this.Info(err)
			} else {
				for _, check := range ei.Usergroup {
					var s models.SunnyUserAndGroup
					s.UserGroupId = check
					s.UserId = int(id)
					models.AddSunnyUserAndGroup(&s)
				}

			}
			this.Success("成功了", "/editor/list", 4)
		}
		//		models.GetSunnyEditorByUsername()
	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		usergroup, _ := models.GetAllUserGroup()
		this.Data["Usergroup"] = usergroup
		this.TplName = "editor/form.html"
	}

}
