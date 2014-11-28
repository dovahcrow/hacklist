package controllers

import (
	"hacklist/controllers/lib"
)

type AboutHandle struct {
	lib.BaseHandle
}

func (this *AboutHandle) Prepare() {
	this.BaseHandle.Prepare()
	this.Layout = "bare.html"
}
func (this *AboutHandle) Get() {
	this.Data[`barTitle`] = "关于我们"
	this.Data[`barDesc`] = `习科道展网络安全顾问`
	this.TplNames = "about.html"
}
