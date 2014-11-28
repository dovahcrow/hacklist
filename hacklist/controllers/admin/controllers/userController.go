package controllers

import (
	"github.com/astaxie/beego"
	"hacklist/controllers/helper"
	"hacklist/models"
	"labix.org/v2/mgo/bson"
)

var numPerPage = 10

type AUserController struct {
	AdminController
}

func (this *AUserController) Prepare() {
	this.AdminController.Prepare()
	this.Data[`position`] = "auser"
	this.Data[`subMenu`] =
		helper.NewSubMenu().
			Set(`用户列表`, this.UrlFor(`AUserController.List`)).
			Set(`添加用户`, this.UrlFor(`AUserController.Add`))
}

func (this *AUserController) List() {
	// fmt.Println(this.UrlFor(`AUserController.Delet
	num, err := models.Usr.GetUserNum()
	if err != nil {
		beego.BeeLogger.Critical("get user count error: %v", err)
		this.Abort("500")
		return
	}
	paginator := helper.NewPaginator(this.Ctx.Request, numPerPage, num)

	users, err := models.Usr.GetUserRange(paginator.Offset(), paginator.PerPageNums)
	if err != nil {
		beego.BeeLogger.Critical("get user count error: %v", err)
		this.Abort("500")
		return
	}

	this.Data[`users`] = users
	this.Data[`paginator`] = paginator
	this.TplNames = "admin/user/list.html"
	return
}
func (this *AUserController) Add() {
	this.TplNames = "index.html"
}
func (this *AUserController) Del() {
	this.Redirect(this.UrlFor("AUserController.List"), 302)
}
func (this *AUserController) Mod() {
	strct := new(models.User)
	err := this.ParseForm(strct)
	if err != nil {
		beego.Warn("mod user:parse req to struct error:", err)
	}
	id := this.Ctx.Input.Param(":id")
	strct.Id = bson.ObjectIdHex(id)

	models.UserCollection.UpdateId(strct.Id.Hex(), update)
}
