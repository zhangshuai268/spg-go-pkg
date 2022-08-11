package aliyun

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestSms_SmsSend(t *testing.T) {
	content := `{"code":"999999"}`
	send, err := client.SmsSend("17685752864", content)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(send.Code)
	return
}
