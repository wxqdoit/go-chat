package db

import (
	"github.com/go-xorm/xorm"
	"log"
)

func InitDb() {
	driverName := "mysql"
	dsName := "root:root@(127.0.0.1:3306)/go_chat"
	DbEngine, err := xorm.NewEngineWithParams(driverName, dsName, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)      //显示sql语句
	DbEngine.SetMaxOpenConns(2) //设置数据库最大连接数
	//DbEngine.Sync2()
}
