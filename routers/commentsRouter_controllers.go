package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:DtfController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/temis_monto_aceptado_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
