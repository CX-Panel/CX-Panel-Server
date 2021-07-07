package main

import (
	_ "CX-Panel-Server/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/", "public")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
