package routers

import (
	"github.com/smarkm/intro-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Include(&controllers.HelloController{})
}
