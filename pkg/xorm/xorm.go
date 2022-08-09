package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Engine struct {
	Orm *xorm.Engine
}

// InitXormEngine 初始化xorm引擎
//  driverName: mysql
//  dataSourceName: user:password@(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=true&loc=Local&collation=utf8mb4_unicode_ci
func InitXormEngine(driverName string, dataSourceName string) (*Engine, error) {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Engine{
		Orm: engine,
	}, nil
}
