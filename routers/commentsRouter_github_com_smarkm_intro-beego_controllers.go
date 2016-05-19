package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/smarkm/intro-beego/controllers:HelloController"] = append(beego.GlobalControllerRouter["github.com/smarkm/intro-beego/controllers:HelloController"],
		beego.ControllerComments{
			"Hello",
			`/hello/:key`,
			[]string{"get"},
			nil})

}
