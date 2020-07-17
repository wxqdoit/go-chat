package db

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"goChat/model"
	"log"
)

var Engine *xorm.Engine

func InitDb() {
	driverName := "mysql"
	dsName := "root:@(127.0.0.1:3306)/go_chat"
	err := errors.New("")
	Engine, err = xorm.NewEngine(driverName, dsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	} else {
		Engine.ShowSQL(true)
		Engine.SetMaxOpenConns(2) //设置数据库最大连接数
		_ = Engine.Sync2(new(model.User))
		println("init ok")
	}
}
