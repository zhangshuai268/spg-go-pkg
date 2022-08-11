package myTime

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_, err := logger.InitLogger(false)
	if err != nil {
		return
	}
	os.Exit(m.Run())
}

func TestStringToMyTime(t *testing.T) {
	s := "2007-12-11 12:32:33"
	time, err := StringToMyTime(s)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	logger.Logger.Info(time)
}
