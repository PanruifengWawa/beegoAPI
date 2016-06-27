package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["newapi/controllers:TTokenController"] = append(beego.GlobalControllerRouter["newapi/controllers:TTokenController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TTokenController"] = append(beego.GlobalControllerRouter["newapi/controllers:TTokenController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TTokenController"] = append(beego.GlobalControllerRouter["newapi/controllers:TTokenController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TTokenController"] = append(beego.GlobalControllerRouter["newapi/controllers:TTokenController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TTokenController"] = append(beego.GlobalControllerRouter["newapi/controllers:TTokenController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TUserController"] = append(beego.GlobalControllerRouter["newapi/controllers:TUserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TUserController"] = append(beego.GlobalControllerRouter["newapi/controllers:TUserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TUserController"] = append(beego.GlobalControllerRouter["newapi/controllers:TUserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TUserController"] = append(beego.GlobalControllerRouter["newapi/controllers:TUserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["newapi/controllers:TUserController"] = append(beego.GlobalControllerRouter["newapi/controllers:TUserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
