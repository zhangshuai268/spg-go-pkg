package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

/**
 * 可用于 结构体转map，map转结构体，json转结构体，结构体数据转移
 * @author zhangshuai
 * @description //TODO
 * @date 15:17 2021/4/13
 * @param
 * @return
 **/
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

/**
 * Get请求
 * @author zhangshuai
 * @description //TODO
 * @date 15:17 2021/4/13
 * @param
 * @return
 */
func HttpGet(url string) (map[string]interface{}, error) {
	var data map[string]interface{}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
