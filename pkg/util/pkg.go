package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"reflect"
	"time"
)

const (
	UpChar  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowChar = "abcdefghijklmnopqrstuvwxyz"
	NumChar = "0123456789"
)

// StructTo 可用于结构体转map，map转结构体，json转结构体，结构体数据转移
func StructTo(old interface{}, new interface{}) error {
	tp := reflect.TypeOf(old)
	var bt []byte
	var err error
	if tp.Name() == "string" {
		bt = []byte(old.(string))
	} else {
		bt, err = json.Marshal(old)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(bt, new)
	if err != nil {
		return err
	}
	return nil
}

// GetRandNum 获取随机字符串
//  n: 生成字符串的位数
//  char: 包含的字符类型，可自定义，可填写多个，自带：UpChar：全部大写字母、LowChar全部小写字母、NumChar 数字字符
func GetRandNum(n int, char ...string) string {
	str := ""
	for _, s := range char {
		str += s
	}
	by := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, by[r.Intn(len(by))])
	}
	return string(result)
}

// Md5Make MD5加密
func Md5Make(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}
