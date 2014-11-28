package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	_ "hacklist/routers"
)

func main() {
	beego.AddFuncMap(`append`, func(i interface{}, t interface{}) string {
		if v, ok := i.(**[]interface{}); ok {
			k := append(**v, t)
			*v = &k
		}
		return ``
	})
	beego.Run()
}
