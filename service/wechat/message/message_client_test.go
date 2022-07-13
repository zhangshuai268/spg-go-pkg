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
	appId       = ""
	appSecret   = ""
	appIdJs     = ""
	appSecretJs = ""
)

func TestMain(m *testing.M) {
	//初始化登录客户端
	client, _ = NewMessageClient(appId, appSecret, wechat.AppTypeO)
	clientJs, _ = NewMessageClient(appIdJs, appSecretJs, wechat.AppTypeA)
	//初始花日志系统
	_, _ = logger.InitLogger(false)

	os.Exit(m.Run())
}
