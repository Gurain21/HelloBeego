package controllers

import (
	"HeeloBeego/db_mysql"
	"HeeloBeego/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	r.Ctx.WriteString("Post请求方式： http://127.0.0.1:8081/register")
	body := r.Ctx.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		r.Ctx.WriteString("body解析错误，请重试！")
		fmt.Println(err.Error())

		return
	}
	var user models.User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		r.Ctx.WriteString("json解析错误，请重试！")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("name",user.Name)
	fmt.Println("birthday:",user.Birtday)
	fmt.Println("address",user.Address)
	fmt.Println("nick",user.Nick)
	rows, err := db_mysql.Inseret(user)
	if err != nil {
		r.Ctx.WriteString("数据库操作错误，请重试！")

		fmt.Println(err.Error())

		return
	}
	fmt.Println("test表中数据被影响的行数：",rows)
}
