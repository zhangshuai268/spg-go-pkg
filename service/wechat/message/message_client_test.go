package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
	"os"
	"testing"
)

var (
	client      MessageService
	clientJs    MessageService
	appId       = "wxd85b4cd83dba5acc"
	appSecret   = "fbae05ceb4a973152bf94ba710a50f6c"
	appIdJs     = "wx091994d55c96a515"
	appSecretJs = "aa8807942087152a94671682bf6ff1eb"
)

func TestMain(m *testing.M) {
	//初始化登录客户端
	client, _ = NewMessageClient(appId, appSecret, wechat.AppTypeO)
	clientJs, _ = NewMessageClient(appIdJs, appSecretJs, wechat.AppTypeA)
	//初始花日志系统
	_, _ = logger.InitLogger(false)

	os.Exit(m.Run())
}
