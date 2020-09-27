package main

import (
	"HeeloBeego/db_mysql"
	_ "HeeloBeego/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/msyql"
)

func main() {
	db_mysql.OpenDB()

	defer db_mysql.Db.Close()

	beego.Run()
}

