package main

import (
	"github.com/astaxie/beego"
	_ "github.com/sunnygocms/managementCMS/routers"
)

func main() {
	//	beego.SetStaticPath("/", "static")
	beego.Run()
}
