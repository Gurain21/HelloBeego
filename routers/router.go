package routers

import (
	"HeeloBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
    beego.Router("/register", &controllers.RegisterController{})

}
