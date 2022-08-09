package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Engine struct {
	Orm *gorm.DB
}

// InitGormEngine 初始化Gorm引擎
//  dsn：user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func InitGormEngine(dsn string) (*Engine, error) {
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Engine{
		Orm: open,
	}, nil
}
