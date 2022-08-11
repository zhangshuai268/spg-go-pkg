package xorm

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestInitXormEngine(t *testing.T) {
	_, err := InitXormEngine("mysql", "root:spg@1703@(121.37.187.166:3306)/yz_test?charset=utf8mb4&parseTime=true&loc=Local&collation=utf8mb4_unicode_ci")
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}
