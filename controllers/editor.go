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
	this.Mapping("List", this.List)
	this.Mapping("Edit", this.Edit)
	this.Mapping("Add", this.Add)
	this.Mapping("Del", this.Del)
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
	//	this.Data["Id"] = this.GetEditorId()
	this.TplName = "editor/list.html"
}

//编辑
func (this *EditorController) Edit() {
	if this.IsSubmit() {
		//此处开始完成保存，保存分两部分一部分是 sunny_editor表，基本完成
		//一部分是sunny_user_and_group表
		ei := EditorInput{}
		if err := this.ParseForm(&ei); err != nil {
			this.Info(err)
		} else {
			if len(ei.Repwd) == 0 || len(ei.Password) == 0 {
				this.Error("密码不能够为空", "-1", 4)
				return
			} else if ei.Password != ei.Repwd {
				this.Error("新的密码两次输入不同", "-1", 4)
				return
			}
			var sunnyeditor models.SunnyEditor
			sunnyeditor.Id = ei.Id
			sunnyeditor.Password = this.SunnyMd5(ei.Password)
			sunnyeditor.Description = ei.Description
			err := models.UpdateSunnyEditorById(&sunnyeditor)
			if err != nil {
				this.Info(err)
			} else {
				//TODO 先删除UserAndGroup里面user_id是这个id的所有组，在重新插入
				//删除权限缓存
				models.DeleteSunnyUserAndGroup(ei.Id)
				for _, check := range ei.Usergroup {
					var s models.SunnyUserAndGroup
					s.UserGroupId = check
					s.UserId = ei.Id
					models.AddSunnyUserAndGroup(&s)
				}
				models.ClearPowerCacheById(ei.Id) //清除权利缓存
			}
			this.Success("成功了", "/editor/list", 4)
		}
	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		usergroup, _ := models.GetAllUserGroup("where active=1")
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
			if len(ei.Username) == 0 {
				this.Error("用户名不能够为空", "-1", 4)
				return
			}
			if len(ei.Repwd) == 0 || len(ei.Password) == 0 {
				this.Error("密码不能够为空", "-1", 4)
				return
			} else if ei.Password != ei.Repwd {
				this.Error("新的密码两次输入不同", "-1", 4)
				return
			}
			//此处判断是否用户名重复
			bEditor := models.IsExistEditorByUsername(ei.Username)
			this.Info(bEditor)
			if bEditor {
				this.Error("这个用户名已经被使用", "-1", 4)
				return
			} else {
				//此处开始完成保存，保存分两部分一部分是 sunny_editor表，基本完成
				//一部分是sunny_user_and_group表
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

		}

	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		usergroup, _ := models.GetAllUserGroup("where active=1")
		this.Data["Usergroup"] = usergroup
		this.TplName = "editor/form.html"
	}

}

//删除用户
func (this *EditorController) Del() {
	var id int
	var sunnyeditor models.SunnyEditor
	mapp := this.Ctx.Input.Params()
	id, _ = strconv.Atoi(mapp["1"])

	sunnyeditor.Id = id
	sunnyeditor.Status = 0
	if err := models.UpdateDelSunnyEditorById(&sunnyeditor); err == nil {
		this.Success("删除成功了", "/editor/list", 4)
	} else {
		this.Error("删除失败", "/editor/list", 4)
	}

}
