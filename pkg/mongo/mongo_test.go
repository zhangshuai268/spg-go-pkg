package mongo

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

func TestInitMongoEngin(t *testing.T) {
	_, err := InitMongoEngin("", "")
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
}
