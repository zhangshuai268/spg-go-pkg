package aliyun

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"os"
	"testing"
)

var (
	client          SmsService
	accessKeyId     = ""
	accessKeySecret = ""
	regionId        = ""
	templateCode    = ""
	signName        = ""
)

func TestMain(m *testing.M) {
	//初始化短信发送客户端
	client, _ = NewSmsClient(accessKeyId, accessKeySecret, regionId, templateCode, signName)
	//初始花日志系统
	_, _ = logger.InitLogger(false)

	os.Exit(m.Run())
}
