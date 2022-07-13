# spg-go-pkg

## 项目介绍
此项目为spg项目组开发封装的相关工具包，为简化开发时相关处理代码，因此封装此包，它将会经常更新。

## 项目下载

`go get github.com/zhangshuai268/spg-go-pkg`

## 项目模块使用

### 一、微信模块(包名:wechat)

#### 1.微信授权登录

##### 使用方法

````
package

import 	wechat "github.com/zhangshuai268/spg-go-pkg/service/wechat/login"

func main() {

    //初始化微信授权client
    client, err := wechat.NewLoginClient("appid", "appsecet")


}

````
##### 接口目录

* [完成接口文档](./doc/wechat/login.md)


#### 2.微信消息发送

##### 使用方法

````

````