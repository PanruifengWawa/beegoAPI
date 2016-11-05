package main

import (
	_ "beegoAPI/docs"
	_ "beegoAPI/routers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/godb")

	//	o := orm.NewOrm()
	//	dr := o.Driver()
	//	fmt.Println(dr.Name())
	//	fmt.Println(dr.Type())
	orm.Debug = true

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)

}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
