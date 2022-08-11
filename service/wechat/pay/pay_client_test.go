package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"os"
	"testing"
)

var (
	appId     = "asa"
	mchId     = "asa"
	paySecret = "asa"
	client    PayService
)

func TestMain(m *testing.M) {
	client, _ = NewPayClient(appId, mchId, paySecret)
	//初始花日志系统
	_, _ = logger.InitLogger(false)

	os.Exit(m.Run())
}
