package redis

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
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

func TestInitRedisEngine(t *testing.T) {
	_, err := InitRedisEngine("127.0.0.1:6379", "", 0)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
}
