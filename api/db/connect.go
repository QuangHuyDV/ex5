package db

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Connect() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", "trung:1234@/ex5")
	if err != nil {
		return engine, err
	}
	engine.ShowSQL(true)
	return engine, err
}
