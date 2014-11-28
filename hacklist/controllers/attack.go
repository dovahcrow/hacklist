package controllers

import (
	"hacklist/controllers/helper"
	"hacklist/controllers/lib"
)

type AttackHandle struct {
	lib.BaseHandle
}

func (this *AttackHandle) Prepare() {
	this.BaseHandle.Prepare()

	this.Data[`subMenu`] = helper.NewSubMenu().
		Set(`所有攻击`, `#`).
		Set(`攻击热点`, `#`).
		Set(`地区统计`, `#`)
	this.Data[`barTitle`] = "攻击统计"
	this.Data[`barDesc`] = "这儿汇总了所有攻击"
	this.Data[`position`] = `attack`
}
func (this *AttackHandle) Get() {
	this.TplNames = "index.html"
}
func (this *AttackHandle) All() {

}
