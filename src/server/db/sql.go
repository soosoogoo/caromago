package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:WeiXin123@/test?charset=utf8")
	if err != nil {
		println(err.Error())
		return
	}

	engine.Alias("o").Where("o.name = 1").Get(&sc)
}
