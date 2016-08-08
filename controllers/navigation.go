package controllers

import (
	//	"errors"
	//	"fmt"
	"strconv"
	"strings"

	"github.com/sunnygocms/managementCMS/models"
	"github.com/sunnygocms/managementCMS/util"
)

type NavigationController struct {
	BaseController
}
type NavigationInput struct {
	Id         int    `form:"id"`
	Name       string `form:"name"`
	Controller string `form:"controller"`
	Action     string `form:"action"`
	Sort       int    `form:"sort"`
	ParentId   int    `form:"parent_id"`
	Display    int8   `form:"display"`
}

func (this *NavigationController) URLMapping() {
	this.Mapping("List", this.List)
	this.Mapping("Edit", this.Edit)
	this.Mapping("Add", this.Add)
	this.Mapping("Del", this.Del)
}
func (this *NavigationController) List() {
	util.ChainClear()
	l, err := models.GetNavigationAll()
	if err != nil {
		this.Data["Navigation"] = err.Error()
	} else {
		for _, sn := range l {
			util.Insert(sn)
		}
		//	util.GetChainLength()
		data, _ := util.ChainToSunnyNavigation()
		//		this.Info(data)
		this.Data["Navigation"] = data
	}

	//	data := l.(models.SunnyNavigation)

	this.TplName = "navigation/list.html"
}
func (this *NavigationController) Edit() {
	if this.IsSubmit() {
		navigationI := NavigationInput{}
		var data models.SunnyNavigation
		if err := this.ParseForm(&navigationI); err != nil {
			this.Info(err)
			this.Error("失败", "-1", 4)
			return
		} else {
			//			this.Info(navigationI)
			if navigationI.Id == navigationI.ParentId {
				this.Error("自己不能是自己的父级菜单", "-1", 4)
				return
			}
			data.Id = navigationI.Id
			data.Name = navigationI.Name
			data.Action = navigationI.Action
			data.Controller = navigationI.Controller
			data.Display = navigationI.Display
			data.Sort = navigationI.Sort
			if navigationI.ParentId == 0 {
				data.ParentId = 0
				data.Level = 0
			} else {
				parent, e := models.GetSunnyNavigationById(navigationI.ParentId)
				if e == nil {
					data.ParentId = navigationI.ParentId
					data.Level = parent.Level + 1
				}
			}
			err := models.UpdateSunnyNavigationById(&data)
			if err != nil {
				this.Error("修改失败", "-1", 4)
				return
			} else {
				this.Success("修改成功了", "/navigation/list", 4)
			}
		}
	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		var id int
		mapp := this.Ctx.Input.Params()
		id, _ = strconv.Atoi(mapp["1"])
		data, _ := models.GetSunnyNavigationById(id)
		this.Data["Data"] = data
		this.TplName = "navigation/form.html"
	}
}
func (this *NavigationController) Add() {
	if this.IsSubmit() {
		navigationI := NavigationInput{}
		var data models.SunnyNavigation
		if err := this.ParseForm(&navigationI); err != nil {
			this.Info(err)
			this.Error("失败", "-1", 4)
			return
		} else {
			//			data.Id = navigationI.Id
			data.Name = navigationI.Name
			data.Action = navigationI.Action
			data.Controller = navigationI.Controller
			data.Display = navigationI.Display
			data.Sort = navigationI.Sort
			if navigationI.ParentId == 0 {
				data.ParentId = 0
				data.Level = 0
			} else {
				parent, e := models.GetSunnyNavigationById(navigationI.ParentId)
				if e == nil {
					data.ParentId = navigationI.ParentId
					data.Level = parent.Level + 1
				}
			}
			_, err := models.AddSunnyNavigation(&data)
			if err != nil {
				this.Error("添加失败", "-1", 4)
				return
			} else {
				this.Success("添加成功了", "/navigation/list", 4)
			}
		}
	} else {
		_, action_name := this.GetControllerAndAction()
		this.Data["ACTION_NAME"] = action_name
		this.TplName = "navigation/form.html"
	}
}
func (this *NavigationController) Del() {
	var id int
	mapp := this.Ctx.Input.Params()
	id, _ = strconv.Atoi(mapp["1"])
	if models.IsExistNavigationSon(id) {
		this.Error("本项目下有子项，不可以删除", "-1", 4)
		return
	}
	if err := models.DeleteSunnyNavigation(id); err == nil {
		this.Success("删除成功了", "/navigation/list", 4)
		return
	} else {
		this.Error("删除失败", "-1", 4)
		return
	}
}

func (this *NavigationController) Jstree() {
	//	this.EnableRender = false
	this.Ctx.Output.Header("Content-Type", "application/javascript; charset=utf-8")
	mapp := this.Ctx.Input.Params()
	//	var data []interface{}
	data, err := models.GetNavigationJsItem()
	if err != nil {
		this.Info(err)
	} else {
		this.Data["json"] = map[string]interface{}{"Items": data}
		//		this.Info(this.SunnyJSON())

		datamap := map[string]string{"Id": mapp["1"], "JustLeaf": mapp["3"], "Value": mapp["5"]}
		s := strings.Replace(this.SunnyJSON(), "\n", "", -1)
		s = strings.Replace(s, " ", "", -1)
		s = strings.Replace(s, `"Id"`, `"id"`, -1)
		s = strings.Replace(s, `"ParentId"`, `"parent_id"`, -1)
		s = strings.Replace(s, `"Name"`, `"name"`, -1)
		this.Data["Jspower"] = s
		this.Data["Data"] = datamap
	}

	this.TplName = "navigation/jstree.html"

}
