package util

import (
	"bytes"
	//	"fmt"

	"github.com/sunnygocms/managementCMS/models"
)

func Hello(in string, o string) (out string) {
	out = in + "world" + o
	return
}

func CheckIsHref(url string, target string, isPower bool) (link string) {
	var buf bytes.Buffer
	if !isPower {
		link = ""
	} else {
		buf.WriteString("<a href='")
		buf.WriteString(url)
		buf.WriteString("'>")
		buf.WriteString(target)
		buf.WriteString("</a>")
		link = buf.String()
	}
	return
}

//这个函数的作用是输入一个字符串根据power判断他是不是可以拥有权限
//有权限返回true

//权限检查
func CheckPower(controller string, action string, id int) (result bool) {
	if id == 1 { //if admin,so have super power
		result = true
	} else {
		power, err := models.GetEditorPowersById(id)
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
	}
	return
}
