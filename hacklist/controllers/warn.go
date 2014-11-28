package controllers

import (
	"hacklist/controllers/helper"
	"hacklist/controllers/lib"
)

type WarnHandle struct {
	lib.BaseHandle
}

func (this *WarnHandle) Prepare() {
	this.BaseHandle.Prepare()

	this.Data[`subMenu`] = helper.NewSubMenu().
		Set(`所有攻击`, `#`)
	this.Data[`barTitle`] = "攻击预警"
	this.Data[`barDesc`] = "预测攻击"
	this.Data[`position`] = `warn`
}
func (this *WarnHandle) Get() {
	this.TplNames = "index.html"
}
