package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestLogin_GetAccessToken(t *testing.T) {
	token, err := client.GetAccessToken()
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	accessToken, err := clientJs.GetAccessToken()
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info(*token)
	logger.Logger.Info(*accessToken)
}
