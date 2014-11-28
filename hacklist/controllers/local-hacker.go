package controllers

import (
	"hacklist/controllers/helper"
	"hacklist/controllers/lib"
)

type LocalHackerHandle struct {
	lib.BaseHandle
}

func (this *LocalHackerHandle) Prepare() {
	this.BaseHandle.Prepare()
	this.Data[`subMenu`] = helper.NewSubMenu().
		Set(`热门黑客`, `#`).
		Set(`所有黑客`, `#`).
		Set(`黑客查询`, `#`)
	this.Data[`barTitle`] = "黑客名册"
	this.Data[`barDesc`] = "在这儿你可以看到一些你熟悉的黑客名字"
	this.Data[`position`] = `hacker`
}
func (this *LocalHackerHandle) Get() {
	this.TplNames = "index.html"
}
