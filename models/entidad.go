package models

type Entidad struct {
	Id_RENAME         int         `orm:"column(id);auto"`
	Nit               string      `orm:"column(nit)"`
	Nombre            string      `orm:"column(nombre)"`
	RolEntidadId      *RolEntidad `orm:"column(rol_entidad_id);rel(fk)"`
	EntidadSucesoraId int         `orm:"column(entidad_sucesora_id);null"`
	EstadoEntidad     string      `orm:"column(estado_entidad)"`
	Direccion         string      `orm:"column(direccion);null"`
	Departament       string      `orm:"column(departament);null"`
	Ciudad            string      `orm:"column(ciudad);null"`
	CorreoElectronico string      `orm:"column(correo_electronico);null"`
	Telefono          string      `orm:"column(telefono);null"`
	Facebook          string      `orm:"column(facebook);null"`
	Twitter           string      `orm:"column(twitter);null"`
	SitioWeb          string      `orm:"column(sitio_web);null"`
}
