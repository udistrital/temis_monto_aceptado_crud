// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/temis_monto_aceptado_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/registrar_monto_aceptado_por_cobrar",
			beego.NSInclude(
				&controllers.RegistrarMontoAceptadoPorCobrarController{},
			),
		),

		beego.NSNamespace("/registrar_recaudo",
			beego.NSInclude(
				&controllers.RegistrarRecaudoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
