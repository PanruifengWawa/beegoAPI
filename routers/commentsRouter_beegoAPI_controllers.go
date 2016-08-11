package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TFkUserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TTokenController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"] = append(beego.GlobalControllerRouter["beegoAPI/controllers:TUserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
