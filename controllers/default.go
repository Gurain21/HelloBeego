package controllers

import (
	"HeeloBeego/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1843607154@qq.com"
	c.TplName = "index.tpl"
	//name1 := c.GetString("name")
	//age1,err := c.GetInt("age")
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
}

//该post 方法时处理post类型的请求时要调用的方法
func (c *MainController)Post() {
	fmt.Println("post类型的请求")
	// Ctx  *context.Context 环境
	user := c.Ctx.Request.Form.Get("user")
	pwd := c.Ctx.Request.Form.Get("pwd")
	//老师的获取From中value的代码
	//user := c.Ctx.Request.FormValue("user")
	//pwd := c.Ctx.Request.FormValue("pwd")
	fmt.Println(user,pwd)
	if user !="wangergou"||pwd !="123456"{
		c.Ctx.ResponseWriter.Write([]byte("sorry,your datas is false!"))
	}else{
		c.Ctx.ResponseWriter.Write([]byte("Welcome, administrator!"))

	}

	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1843607154@qq.com"
	c.TplName = "index.tpl"

	//body := c.Ctx.Request.Body
	bytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		fmt.Println(err.Error())
		c.Ctx.WriteString("数据接收失败")
	}
	//var user1 models.Person
	//err = json.Unmarshal(bytes,&user1)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	c.Ctx.ResponseWriter.Write([]byte("数据解析失败"))
	//}
	//fmt.Println("用户名：",user1.Name)
	//fmt.Println("用户年龄：",user1.Age)
	//fmt.Println("用户性别：",user1.Sex)
	var person1 models.Personal
	err = json.Unmarshal(bytes,&person1)
	if err != nil {
		fmt.Println(err.Error())
		c.Ctx.WriteString("数据解析失败。。。。")

	}
	fmt.Println("姓名",person1.Name)
	fmt.Println("生日",person1.Birthday)
	fmt.Println("地址",person1.Address)
	fmt.Println("简介",person1.Nick)
}
