// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/oikos/oikos_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaController{},
			),
		),

		beego.NSNamespace("/dependencia_padre",
			beego.NSInclude(
				&controllers.DependenciaPadreController{},
			),
		),

		beego.NSNamespace("/tipo_espacio_fisico",
			beego.NSInclude(
				&controllers.TipoEspacioFisicoController{},
			),
		),

		beego.NSNamespace("/asignacion_espacio_fisico_dependencia",
			beego.NSInclude(
				&controllers.AsignacionEspacioFisicoDependenciaController{},
			),
		),

		beego.NSNamespace("/espacio_fisico",
			beego.NSInclude(
				&controllers.EspacioFisicoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
