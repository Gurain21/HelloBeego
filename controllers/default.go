package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1843607154@qq.com"
	c.TplName = "index.tpl"

	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
}

//该post 方法时处理post类型的请求时要调用的方法
func (c *MainController)Post() {
	fmt.Println("post类型的请求")
	// Ctx  *context.Context
	user := c.Ctx.Request.Form.Get("user")
	pwd := c.Ctx.Request.Form.Get("pwd")
	//老师的获取From中value的代码
	//user := c.Ctx.Request.FormValue("user")
	//pwd := c.Ctx.Request.FormValue("pwd")
	fmt.Println(user,pwd)
	if user !="周智杰"||pwd !="123456"{
		c.Ctx.ResponseWriter.Write([]byte("sorry,your datas is false!"))
	}else{
		c.Ctx.ResponseWriter.Write([]byte("Welcome, administrator!"))

	}




	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1843607154@qq.com"
	c.TplName = "index.tpl"
}
