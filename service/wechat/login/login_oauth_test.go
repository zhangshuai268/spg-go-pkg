package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestLogin_Oauth(t *testing.T) {
	code := "083lBDkl2Og5w94Ttyol25QyZh2lBDka"
	oauth, err := client.Oauth(code)
	if err != nil {
		logger.Logger.Info(err.Error())
		return
	}
	logger.Logger.Info(oauth.OpenId)
	logger.Logger.Info(oauth.AccessToken)
}
