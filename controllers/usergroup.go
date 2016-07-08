package controllers

import (
	"strconv"

	"github.com/sunnygocms/managementCMS/models"
)

// oprations for SunnyEditor
type UsergroupController struct {
	BaseController
}
type UsergroupInput struct {
	Id          int    `form:"id"`
	GroupName   string `form:"group_name"`
	Active      int    `form:"active"`
	Power       []int  `form:"power"`
	Description string `form:"description"`
}

func (this *UsergroupController) URLMapping() {
	this.Mapping("List", this.List)
	this.Mapping("Edit", this.Edit)
	this.Mapping("Add", this.Add)
	this.Mapping("Del", this.Del)
}
func (this *UsergroupController) List() {
	l, err := models.GetAllUserGroup("")
	if err != nil {
		this.Data["Usergroup"] = err.Error()
	} else {
		this.Data["Usergroup"] = l
	}
	models.ClearPowerCacheAllById()
	this.Data["Id"] = this.GetEditorId()
	this.TplName = "usergroup/list.html"
}

//添加
func (this *UsergroupController) Add() {
	if this.IsSubmit() {
		ui := UsergroupInput{}
		if err := this.ParseForm(&ui); err != nil {
			this.Info(err)
			this.Error("失败", "-1", 4)
			return
		} else {
			if len(ui.GroupName) == 0 {
				this.Error("用户组名不能够为空", "-1", 4)
				return
			}
			//			this.Info(ui)
			//此处判断是否用户名重复
			bEditor := models.IsExistUsergroupByUsername(ui.GroupName)
			//			this.Info(bEditor)
			if bEditor {
				this.Error("这个用户名已经被使用", "-1", 4)
				return
			}

			var data models.SunnyUserGroup
			data.GroupName = ui.GroupName
			data.Description = ui.Description
			data.EditId = this.GetEditorId()
			if ui.Active == 1 {
				data.Active = 1
			} else {
				data.Active = 0
			}

			id, err := models.AddSunnyUserGroup(&data)
			if err != nil {
				this.Info(err)
			} else {
				models.DeleteSunnyUsergroupAndPower(int(id)) //删除这个序号的所有权限
				if ui.Active == 1 {                          //只有可用的才可以添加权限
					for _, check := range ui.Power {
						var s models.SunnyUsergroupAndPower
						s.PowerId = check
						s.UserGroupId = int(id)
						models.AddSunnyUsergroupAndPower(&s)
					}
				}
				this.Success("成功了", "/usergroup/list", 4)
			}
		}
	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		power, _ := models.GetSunnyPowerAll()
		this.Data["json"] = power
		this.Data["Jspower"] = this.SunnyJSON()
		this.Data["Power"] = power
		this.Data["UserGroupId"] = 0
		this.TplName = "usergroup/form.html"
	}

}

func (this *UsergroupController) Edit() {
	if this.IsSubmit() {
		ui := UsergroupInput{}
		if err := this.ParseForm(&ui); err != nil {
			this.Info(err)
			this.Error("失败", "-1", 4)
			return
		} else {
			var data models.SunnyUserGroup
			data.Id = ui.Id
			data.GroupName = ui.GroupName
			data.Description = ui.Description
			data.EditId = this.GetEditorId()
			if ui.Active == 1 {
				data.Active = 1
			} else {
				data.Active = 0
			}
			err := models.UpdateSunnyUserGroupById(&data)
			if err != nil {
				this.Success("修改失败", "/usergroup/list", 4)
				return
			} else {
				err := models.DeleteSunnyUsergroupAndPower(ui.Id) //删除这个序号的所有权限
				this.Info(err)
				if ui.Active == 1 { //只有可用的才可以添加权限
					for _, check := range ui.Power {
						var s models.SunnyUsergroupAndPower
						s.PowerId = check
						s.UserGroupId = ui.Id
						models.AddSunnyUsergroupAndPower(&s)
					}
				}
			}
			models.ClearPowerCacheAllById()

			this.Success("修改成功了", "/usergroup/list", 4)
		}
	} else {
		this.Data["UserGroupId"] = 1
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		power, _ := models.GetSunnyPowerAll()
		this.Data["Power"] = power

		var id int
		mapp := this.Ctx.Input.Params()
		id, _ = strconv.Atoi(mapp["1"])
		usergroup, err := models.GetSunnyUserGroupById(id)
		if err == nil {
			this.Data["Data"] = &usergroup
		} else {
			this.Info(err)
		}

		var usergroupandpower []interface{}
		usergroupandpower, err = models.GetSunnyUsergroupAndPowerById(id)
		//		this.Info(usergroupandpower)
		this.Data["json"] = usergroupandpower
		this.Data["Jspower"] = this.SunnyJSON()

		this.TplName = "usergroup/form.html"
	}
}

func (this *UsergroupController) Del() {
	var id int
	mapp := this.Ctx.Input.Params()
	id, _ = strconv.Atoi(mapp["1"])
	models.DeleteSunnyUsergroupAndPower(id)         //删除这个序号的所有权限
	models.DeleteSunnyUserAndGroupByUserGroupId(id) //删除用户属于这个组
	if err := models.DeleteSunnyUserGroup(id); err == nil {
		this.Success("删除成功了", "/usergroup/list", 4)
	} else {
		this.Success("删除失败", "/usergroup/list", 4)
	}
}
