package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sunnygocms/managementCMS/routers"
	"github.com/sunnygocms/managementCMS/util"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:green@)!#ok@tcp(127.0.0.1:3306)/sunny")
	orm.Debug = true
	orm.RunSyncdb("default", false, false)
}
func main() {
	//	beego.SetStaticPath("/", "static")
	beego.AddFuncMap("checkIsHref", util.CheckPower)
	beego.Run()
}
