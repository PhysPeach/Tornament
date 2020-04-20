package routers

import (
	"github.com/PhysPeach/Tornament/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.RootController{})
}
