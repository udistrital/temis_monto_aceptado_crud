package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarMontoAceptadoPorCobrarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"] = append(beego.GlobalControllerRouter["github.com/temis_monto_aceptado_crud/controllers:RegistrarRecaudoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
