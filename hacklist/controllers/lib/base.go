package lib

import (
	"github.com/astaxie/beego"
)

type BaseHandle struct {
	beego.Controller
}

func (this *BaseHandle) Prepare() {
	j := &[]interface{}{}
	c := &[]interface{}{}
	this.Data[`moreStyles`] = &c
	this.Data[`moreScripts`] = &j
	this.Data[`position`] = ``
	this.Layout = `frame.html`
}
