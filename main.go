package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sunnygocms/managementCMS/routers"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "sunnydba:123456@tcp(127.0.0.1:3306)/sunny")
}
func main() {
	//	beego.SetStaticPath("/", "static")
	beego.Run()
}
