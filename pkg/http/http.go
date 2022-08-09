package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// HttpGet Get请求
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

// HttpPostJson post请求通过json
func HttpPostJson(url string, post []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	res, err := http.Post(url, "application/json", bytes.NewReader(post))
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

// HttpPostForm post请求通过form
func HttpPostForm(urls string, post map[string]string) (map[string]interface{}, error) {
	var data map[string]interface{}
	postData := make(url.Values)
	for key, value := range post {
		postData[key] = []string{value}
	}
	res, err := http.PostForm(urls, postData)
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

// GetIP 获取ip
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
