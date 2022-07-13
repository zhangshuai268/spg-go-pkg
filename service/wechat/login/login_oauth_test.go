package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestLogin_Oauth(t *testing.T) {
	code := "043CKLll2n47v94491pl2nVIsi3CKLlr"
	oauth, err := client.Oauth(code)
	if err != nil {
		logger.Logger.Info(err.Error())
		return
	}
	logger.Logger.Info(oauth.OpenId)
	logger.Logger.Info(oauth.AccessToken)
}
