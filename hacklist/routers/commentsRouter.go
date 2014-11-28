package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"] = append(beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"],
		beego.ControllerComments{
			"List",
			"/hacker/list",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"] = append(beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"],
		beego.ControllerComments{
			"Add",
			"/hacker/add",
			[]string{"get","post"},
			nil})

	beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"] = append(beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"],
		beego.ControllerComments{
			"Del",
			"/hacker/delete/:id",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"] = append(beego.GlobalControllerRouter["hacklist/controllers/admin/controllers:AHackerController"],
		beego.ControllerComments{
			"Mod",
			"/hacker/modify/:id",
			[]string{"get"},
			nil})

}
