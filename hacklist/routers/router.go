package routers

import (
	"github.com/astaxie/beego"
	"hacklist/controllers"
	"hacklist/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	beego.Router("/attack", &controllers.AttackHandle{})

	beego.Router("/hacker/local", &controllers.LocalHackerHandle{})
	beego.Router("/hacker/foreign", &controllers.ForeignHackerHandle{})
	beego.Router("/search", &controllers.SearchHandle{})

	beego.Router("/about", &controllers.AboutHandle{})

	beego.Router("/warn", &controllers.WarnHandle{})

	beego.AddNamespace(admin.AdminNS)
}
