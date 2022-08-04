package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestPay_PayAddCertFilePath(t *testing.T) {
	certFilePath := ""
	keyFilePath := ""
	err := client.PayAddCertFilePath(certFilePath, keyFilePath)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
}