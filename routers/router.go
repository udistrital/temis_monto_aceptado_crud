// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/casossat/temis_monto_aceptado_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/rol_entidad",
			beego.NSInclude(
				&controllers.RolEntidadController{},
			),
		),

		beego.NSNamespace("/experiencia_laboral",
			beego.NSInclude(
				&controllers.ExperienciaLaboralController{},
			),
		),

		beego.NSNamespace("/experiencia_laboral_incapacidad",
			beego.NSInclude(
				&controllers.ExperienciaLaboralIncapacidadController{},
			),
		),

		beego.NSNamespace("/incapacidad",
			beego.NSInclude(
				&controllers.IncapacidadController{},
			),
		),

		beego.NSNamespace("/indice_precio_consumidor",
			beego.NSInclude(
				&controllers.IndicePrecioConsumidorController{},
			),
		),

		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/salario_minimo_legal",
			beego.NSInclude(
				&controllers.SalarioMinimoLegalController{},
			),
		),

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

		beego.NSNamespace("/organizacion",
			beego.NSInclude(
				&controllers.OrganizacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
