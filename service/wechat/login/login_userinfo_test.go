package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"github.com/zhangshuai268/spg-go-pkg/service/wechat"
	"testing"
)

func TestLogin_UserInfo(t *testing.T) {
	openId := ""
	accessToken := ""
	info, err := client.UserInfo(accessToken, openId, wechat.LandCN)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(*info)
}
