package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestLogin_JsCode2(t *testing.T) {
	code := "083lBDkl2Og5w94Ttyol25QyZh2lBDka"
	res, err := clientJs.JsCode2(code)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Api(res.OpenId)
	logger.Logger.Api(res.SessionKey)
	logger.Logger.Api(res.ErrMessage)
	return
}