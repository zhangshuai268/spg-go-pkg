package aliyun

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestSms_SmsSendDetail(t *testing.T) {
	res, err := client.SmsSendDetail("", "", 10, 1)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(res.SmsSendDetailDTOs)
	return
}
