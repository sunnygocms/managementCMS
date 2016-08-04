package controllers

import (
	//	"errors"
	//	"fmt"
	//	"strconv"
	//	"strings"

	"github.com/sunnygocms/managementCMS/models"
	"github.com/sunnygocms/managementCMS/util"
)

type NavigationController struct {
	BaseController
}

func (this *NavigationController) URLMapping() {
	this.Mapping("List", this.List)
	this.Mapping("Edit", this.Edit)
	this.Mapping("Add", this.Add)
	this.Mapping("Del", this.Del)
}
func (this *NavigationController) List() {
	l, err := models.GetNavigationAll()
	if err != nil {
		this.Data["Navigation"] = err.Error()
	} else {
		this.Data["Navigation"] = l
	}

	//	data := l.(models.SunnyNavigation)
	for _, sn := range l {
		util.Insert(sn)
	}
	this.Info(util.Chain)
	this.TplName = "navigation/list.html"
}
func (this *NavigationController) Edit() {
	this.Html("Navigation")
}
func (this *NavigationController) Add() {
	this.Html("Navigation")
}
func (this *NavigationController) Del() {
	this.Html("Navigation")
}
