package sdk

import (
	login "github.com/zhangshuai268/spg-go-pkg/service/wx/login/v1"
)

type Client interface {
	WxLogin(param interface{}) (login.LoginService, error)
}

type client struct {
}

func (c *client) WxLogin(param interface{}) (login.LoginService, error) {
	newLogin, err := login.NewLogin(param)
	if err != nil {
		return nil, err
	}
	return newLogin, nil
}

func NewClient() (Client, error) {
	var c *client
	return c, nil
}
