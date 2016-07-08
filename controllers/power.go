package controllers

import (
	//	"errors"
	//			"fmt"
	"strconv"
	//			"strings"

	"github.com/sunnygocms/managementCMS/models"
)

// oprations for SunnyEditor
type PowerController struct {
	BaseController
}
type PowerInput struct {
	Id         int    `form:"_"`
	PowerName  string `form:"power_name"`
	Controller string `form:"controller"`
	Action     string `form:"action"`
}

func (this *PowerController) URLMapping() {
	this.Mapping("List", this.List)
	this.Mapping("Edit", this.Edit)
	this.Mapping("Add", this.Add)
	this.Mapping("Del", this.Del)
}

func (this *PowerController) List() {
	l, err := models.GetPowerAll()
	if err != nil {
		this.Data["Power"] = err.Error()
	} else {
		this.Data["Power"] = l
	}
	this.Info(l)
	this.Data["Power"] = l
	this.Data["Id"] = this.GetEditorId()
	this.TplName = "power/list.html"
}

func (this *PowerController) Add() {
	if this.IsSubmit() {
		powerI := PowerInput{}
		if err := this.ParseForm(&powerI); err != nil {
			this.Info(err)
			this.Error("失败", "-1", 4)
			return
		} else {
			if len(powerI.PowerName) == 0 {
				this.Error("限名称不能够为空", "-1", 4)
				return
			}
			//判断名称是否重复
			bEditor := models.IsExistPowerByUsername(powerI.PowerName)

			if bEditor {
				this.Error("这个权限名称已经被使用", "-1", 4)
				return
			}
			//			models.ClearPowerCacheById(powerI.Id)
		}
	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		this.TplName = "power/form.html"
	}
}

func (this *PowerController) Edit() {
	if this.IsSubmit() {
		powerI := PowerInput{}
		if err := this.ParseForm(&powerI); err != nil {
			this.Info(err)
			this.Error("失败", "-1", 4)
			return
		} else {
			if len(powerI.PowerName) == 0 {
				this.Error("限名称不能够为空", "-1", 4)
				return
			}

			var data models.SunnyPower
			data.Id = powerI.Id
			//			data. = powerI
			//			data.Description = powerI.Description
			//			data.EditId = this.GetEditorId()
			models.ClearPowerCacheById(powerI.Id)
		}

	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		var id int
		mapp := this.Ctx.Input.Params()
		id, _ = strconv.Atoi(mapp["1"])
		data, _ := models.GetSunnyPowerById(id)
		this.Data["Data"] = data
		this.TplName = "power/form.html"
	}

}

func (this *PowerController) Del() {
	models.ClearPowerCacheAll()
}
