package main

import (
	"FEINESV/models"
	_ "FEINESV/routers"
	"FEINESV/spymongo"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
	spymongo.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.SetStaticPath("/admin/static", "static")
	beego.SetStaticPath("/admin/src", "src")
	beego.Run()
}
