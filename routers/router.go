// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegoAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/t_fk_user",
			beego.NSInclude(
				&controllers.TFkUserController{},
			),
		),

		beego.NSNamespace("/t_token",
			beego.NSInclude(
				&controllers.TTokenController{},
			),
		),

		beego.NSNamespace("/t_user",
			beego.NSInclude(
				&controllers.TUserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
