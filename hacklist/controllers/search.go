package controllers

import (
	"hacklist/controllers/helper"
	"hacklist/controllers/lib"
)

type SearchHandle struct {
	lib.BaseHandle
}

func (this *SearchHandle) Prepare() {
	this.BaseHandle.Prepare()

	this.Data[`subMenu`] = helper.NewSubMenu().
		Set(`站点搜索`, `#`)
	this.Data[`barTitle`] = "站点搜索"
	this.Data[`barDesc`] = "看看你的网站有没有被黑"
	this.Data[`position`] = `search`
}
func (this *SearchHandle) Get() {
	this.TplNames = "index.html"
}
