package gorm

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestInitGormEngine(t *testing.T) {
	_, err := InitGormEngine("root:spg@1703@tcp(121.37.187.166:3306)/yz_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
}