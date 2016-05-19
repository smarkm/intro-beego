package controllers

import "github.com/astaxie/beego"

//HelloController is hello test
type HelloController struct {
	beego.Controller
}

// @router /hello/:key [get]
func (c *HelloController) Hello() {
	key := c.GetString("key", "def")
	println(key)
	c.Data["json"] = "hello message from server for "+key
	c.ServeJson()
}
