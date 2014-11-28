package controllers

import (
	"hacklist/controllers/lib"
)

type IndexController struct {
	lib.BaseHandle
}

func (this *IndexController) Prepare() {
	this.BaseHandle.Prepare()
	this.Data[`position`] = `index`
	this.Layout = "bare.html"
}
func (this *IndexController) Get() {
	this.Data[`barTitle`] = "这是主页"
	this.Data[`barDesc`] = "这是描述"
	this.TplNames = "index.html"
}
