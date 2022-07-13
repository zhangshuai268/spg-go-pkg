package wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

// MobileResponse 手机号解析Rsp
type MobileResponse struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	WaterMark       struct {
		AppId     string `json:"appid"`
		TimeStamp int    `json:"timestamp"`
	} `json:"watermark"`
}

func (l *login) Mobile(mobile, iv, sessionKey string) (*MobileResponse, error) {
	s, err := dncrypt(mobile, iv, sessionKey)
	if err != nil {
		return nil, err
	}
	var res MobileResponse
	err = util.StructTo(s, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func dncrypt(rawData, iv, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	key_b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	iv_b, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}

	dnData, err := aesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 解密
func aesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = pKCS7UnPadding(encryptData)
	return encryptData, nil
}

//去除填充
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
