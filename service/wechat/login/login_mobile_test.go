package wechat

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/log"
	"testing"
)

func TestLogin_Mobile(t *testing.T) {
	mobile := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
	session := "lyY4HPQbaOYzZdG+JcYK9w=="
	res, err := client.Mobile(mobile, iv, session)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(*res)
}
