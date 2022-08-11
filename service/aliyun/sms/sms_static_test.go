package aliyun

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"github.com/zhangshuai268/spg-go-pkg/service/aliyun"
	"testing"
)

func TestSms_SmsSendStatic(t *testing.T) {
	static, err := client.SmsSendStatic(aliyun.GlobalIn, "20220804", "20220804", 1, 10)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(static.Data)
	return
}
