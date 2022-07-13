package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"os"
	"testing"
)

var (
	client      LoginService
	clientJs    LoginService
	appId       = "wxd85b4cd83dba5acc"
	appSecret   = "fbae05ceb4a973152bf94ba710a50f6c"
	appIdJs     = "wx091994d55c96a515"
	appSecretJs = "aa8807942087152a94671682bf6ff1eb"
)

func TestMain(m *testing.M) {
	//初始化登录客户端
	client, _ = NewLoginClient(appId, appSecret)
	clientJs, _ = NewLoginClient(appIdJs, appSecretJs)
	//初始花日志系统
	_, _ = logger.InitLogger(false)

	os.Exit(m.Run())
}
