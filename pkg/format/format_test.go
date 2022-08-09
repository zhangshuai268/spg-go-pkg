package format

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_, err := logger.InitLogger(false)
	if err != nil {
		return
	}
	os.Exit(m.Run())
}

func TestVerifyEmailFormat(t *testing.T) {
	format := VerifyEmailFormat("@qq.com")
	emailFormat := VerifyEmailFormat(".com")
	logger.Logger.Info(format)
	logger.Logger.Info(emailFormat)
}

func TestVerifyIdFormat(t *testing.T) {
	format := VerifyIdFormat("")
	idFormat := VerifyIdFormat("")
	logger.Logger.Info(format)
	logger.Logger.Info(idFormat)
}

func TestVerifyMobileFormat(t *testing.T) {
	format := VerifyMobileFormat("")
	mobileFormat := VerifyMobileFormat("")
	logger.Logger.Info(format)
	logger.Logger.Info(mobileFormat)
}
