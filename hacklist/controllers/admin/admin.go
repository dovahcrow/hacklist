package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hacklist/controllers/admin/controllers"
)

var AdminNS = beego.NewNamespace("/admin")

func init() {
	AdminNS.Include(&controllers.AdminController{})
	AdminNS.Include(&controllers.AHackerController{})
	AdminNS.Include(&controllers.AAttackController{})
	AdminNS.Include(&controllers.AUserController{})

	AdminNS.Router(`/login`, &controllers.AdminController{}, "get,post:Login")
	AdminNS.Router(`/status`, &controllers.AdminController{}, "get:Status")
	AdminNS.Router(`/exit`, &controllers.AdminController{}, "get:Exit")
	AdminNS.Get(`/index`,
		func(ctx *context.Context) {
			ctx.Redirect(302, beego.UrlFor(`AdminController.Status`))
			return
		})
	AdminNS.Get(`/`,
		func(ctx *context.Context) {
			ctx.Redirect(302, beego.UrlFor(`AdminController.Status`))
			return
		})
	AdminNS.Router(`/user/modify/:id`, &controllers.AUserController{}, "get,post:Mod")
	AdminNS.Router(`/user/delete/:id`, &controllers.AUserController{}, "get:Del")
	AdminNS.Router(`/user/list`, &controllers.AUserController{}, "get:List")
	AdminNS.Router(`/user/add`, &controllers.AUserController{}, "get,post:Add")

	AdminNS.Router(`/hacker/modify/:id`, &controllers.AHackerController{}, "get,post:Mod")
	AdminNS.Router(`/hacker/delete/:id`, &controllers.AHackerController{}, "get:Del")
	AdminNS.Router(`/hacker/list`, &controllers.AHackerController{}, "get:List")
	AdminNS.Router(`/hacker/add`, &controllers.AHackerController{}, "get:Add")

	AdminNS.Router(`/attack/modify/:id`, &controllers.AAttackController{}, "get,post:Mod")
	AdminNS.Router(`/attack/delete/:id`, &controllers.AAttackController{}, "get:Del")
	AdminNS.Router(`/attack/list`, &controllers.AAttackController{}, "get:List")
	AdminNS.Router(`/attack/add`, &controllers.AAttackController{}, "get:Add")
}
