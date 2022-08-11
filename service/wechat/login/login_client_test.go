package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"os"
	"testing"
)

var (
	client      LoginService
	clientJs    LoginService
	appId       = ""
	appSecret   = ""
	appIdJs     = ""
	appSecretJs = ""
)

func TestMain(m *testing.M) {
	//初始化登录客户端
	client, _ = NewLoginClient(appId, appSecret)
	clientJs, _ = NewLoginClient(appIdJs, appSecretJs)
	//初始花日志系统
	_, _ = logger.InitLogger(false)

	os.Exit(m.Run())
}
