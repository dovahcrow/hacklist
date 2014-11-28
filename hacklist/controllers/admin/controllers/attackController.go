package controllers

import (
	"github.com/astaxie/beego"
	"hacklist/controllers/helper"
	"hacklist/models"
)

type AAttackController struct {
	AdminController
}

func (this *AAttackController) Prepare() {
	this.AdminController.Prepare()
	this.Data[`position`] = "aattack"
	this.Data[`subMenu`] =
		helper.NewSubMenu().
			Set(`攻击列表`, this.UrlFor(`AAttackController.List`)).
			Set(`添加记录`, this.UrlFor(`AAttackController.Add`))
}

func (this *AAttackController) List() {

	num, err := models.Atk.GetAttackNum()
	if err != nil {
		beego.BeeLogger.Critical("get attack count error: %v", err)
		this.Abort("500")
		return
	}
	paginator := helper.NewPaginator(this.Ctx.Request, numPerPage, num)

	attacks, err := models.Atk.GetAttackRange(paginator.Offset(), paginator.PerPageNums)
	if err != nil {
		beego.BeeLogger.Critical("get attacks error: %v", err)
		this.Abort("500")
		return
	}
	this.Data[`attacks`] = attacks
	this.Data[`paginator`] = paginator
	this.TplNames = "admin/attack/list.html"
	return
}

func (this *AAttackController) Add() {
	this.TplNames = "index.html"
}
func (this *AAttackController) Del() {
	this.Redirect(this.UrlFor("AUserController.List"), 302)
}
func (this *AAttackController) Mod() {

}
