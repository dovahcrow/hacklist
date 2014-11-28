package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hacklist/controllers/lib"
	"hacklist/models"
)

type AdminController struct {
	user *models.User
	lib.BaseHandle
}

func (this *AdminController) Prepare() {
	this.BaseHandle.Prepare()
	beego.ReadFromRequest(&this.Controller)
	this.Layout = `admin/frame.html`

	//user control

	if this.Ctx.Request.RequestURI == "/admin/login" {
		return
	}
	var userID string
	if u := this.GetSession(`userID`); u == nil {
		this.Redirect(this.UrlFor("AdminController.Login"), 302)
		return
	} else {
		userID = u.(string)
	}
	err := models.Usr.IncrUserLogin(userID)
	if err != nil {
		beego.BeeLogger.Warn("Update user login time error: %v", err)
	}
	this.user, err = models.Usr.GetUserById(userID)
	if err != nil {
		beego.BeeLogger.Critical("get user %v info fail: %v", userID, err)
	}
	this.Data[`self`] = this.user

}

func (this *AdminController) Login() {
	if this.Ctx.Input.Method() == "GET" {
		this.Data[`barTitle`] = `登陆`
		this.Layout = "admin/bare.html"
		this.TplNames = "admin/login.html"

		return
	} else if this.Ctx.Input.Method() == "POST" {
		username := this.GetString(`username`)
		password := this.GetString(`password`)
		user, err := models.Usr.GetUser(username, password)
		if err != nil {

			beego.BeeLogger.Warn(
				"user login with %v and %v from %v but failed: %v",
				username,
				password,
				this.Ctx.Input.Request.RemoteAddr,
				err,
			)

			flash := beego.NewFlash()
			flash.Error("用户名或密码不正确")
			flash.Store(&this.Controller)
			this.Redirect(this.UrlFor(".Login"), 302)
			return
		}

		this.SetSession("userID", user.Id.Hex())
		this.Redirect(this.UrlFor(".Status"), 302)
		return
	}
}

func (this *AdminController) Status() {
	this.Layout = "admin/bare.html"
	this.Data[`position`] = `astatus`
	this.Data[`barTitle`] = fmt.Sprintf("你好,管理员 %v", this.user.Name)
	this.Data[`barDesc`] = `在这里,你或许可以为所欲为`
	this.TplNames = "admin/index.html"
	return
}

func (this *AdminController) Exit() {
	this.DelSession("userID")
	this.Redirect(this.UrlFor("AdminController.Login"), 302)
	return
}
