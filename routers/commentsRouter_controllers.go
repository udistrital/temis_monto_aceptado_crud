package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:DtfController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:ExperienciaLaboralIncapacidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IncapacidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:IndicePrecioConsumidorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:RolEntidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:SalarioMinimoLegalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/casossat/TEMIS/midApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
