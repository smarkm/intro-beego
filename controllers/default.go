package controllers

import (
	"github.com/astaxie/beego"
)

//MainController commont
type MainController struct {
	beego.Controller
}
//Get commont
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}
