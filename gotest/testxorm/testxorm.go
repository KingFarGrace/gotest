package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type User struct {
	Uid  int `xorm:"'uid' pk autoincr"`
	Name string
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:132546@/testxorm?charset=utf8")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	user := new(User)
	user.Name = "King"
	engine.SetMapper(names.GonicMapper{})
	_, err = engine.Insert(user)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}
