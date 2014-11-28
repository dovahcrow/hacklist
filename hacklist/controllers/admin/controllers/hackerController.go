package controllers

import (
	"github.com/astaxie/beego"
	"hacklist/controllers/helper"
	"hacklist/models"
	"time"
)

type AHackerController struct {
	AdminController
}

func (this *AHackerController) URLMapping() {
	this.Mapping(`List`, this.List)
}
func (this *AHackerController) Prepare() {
	this.AdminController.Prepare()
	this.Data[`position`] = "ahacker"
	this.Data[`barTitle`] = "黑客记录"
	this.Data[`barDesc`] = "在这里编辑黑客记录"
	this.Data[`subMenu`] =
		helper.NewSubMenu().
			Set(`黑客列表`, this.UrlFor(`AHackerController.List`)).
			Set(`添加记录`, this.UrlFor(`AHackerController.Add`))
}

// @router /hacker/list [get]
func (this *AHackerController) List() {
	num, err := models.Hkr.GetHackerNum()
	if err != nil {
		beego.BeeLogger.Critical("get hacker count error: %v", err)
		this.Abort("500")
		return
	}
	paginator := helper.NewPaginator(this.Ctx.Request, numPerPage, num)

	hackers, err := models.Hkr.GetHackerRange(paginator.Offset(), paginator.PerPageNums)
	if err != nil {
		beego.BeeLogger.Critical("get hackers error: %v", err)
		this.Abort("500")
		return
	}
	this.Data[`hackers`] = hackers
	this.Data[`paginator`] = paginator
	this.TplNames = "admin/hacker/list.html"
	return
}

// @router /hacker/add [get,post]
func (this *AHackerController) Add() {
	this.TplNames = "admin/hacker/add.html"
	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		defer flash.Store(&this.Controller)

		h := new(models.Hacker)
		err := this.ParseForm(h)
		if err != nil {
			beego.BeeLogger.Critical("add hacker: parse form fail: %v", err)
			flash.Error("解析表单错误: %v", err)
			return
		}
		flash.Data[`nick`] = h.Nick

		flash.Data[`gender`] = h.Gender

		flash.Data[`realname`] = h.RealName
		flash.Data[`ID`] = h.PersonalIdentify
		birthdayS := this.GetString("birthday")
		flash.Data[`birthday`] = birthdayS
		birthday, err := time.Parse("2006-01-02", birthdayS)
		if err != nil {
			beego.BeeLogger.Critical("add hacker failed. parse birthday %v error: %v", birthdayS, err)
			flash.Error(`解析生日错误`)
			return
		}
		h.Birthday = birthday
		err = models.Hkr.InsertHacker(h)
		if err != nil {
			beego.BeeLogger.Critical("add hacker: insert to db fail: %v", err)
			flash.Error(`内部错误`)
			return
		}
		flash.Notice(`插入成功`)
		for i := range flash.Data {
			if i == "notice" {
				continue
			} else {
				delete(flash.Data, i)
			}
		}
	} else if this.Ctx.Input.Method() == "GET" {
		return
	}

}

// @router /hacker/delete/:id [get]
func (this *AHackerController) Del() {
	id := this.Ctx.Input.Param(`:id`)
	err := models.Hkr.DeleteHackerById(id)
	if err != nil {
		beego.BeeLogger.Critical("Delete hacker %v fail: %v", id, err)
		this.Abort("500")
		return
	}
	flash := beego.NewFlash()
	flash.Notice("删除记录成功")
	flash.Store(&this.Controller)
	this.Redirect(this.UrlFor("AHackerController.List"), 302)
}

// @router /hacker/modify/:id [get]
func (this *AHackerController) Mod() {

}
