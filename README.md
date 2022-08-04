# spg-go-pkg

## 项目介绍

此项目为spg项目组开发封装的相关工具包，为简化开发时相关处理代码，因此封装此包，它将会经常更新。

> #### 本包中所有方法均有参数注释，在调用时请注意相关代码注释或者代码编译工具的提示！

## 项目下载

`go get github.com/zhangshuai268/spg-go-pkg`

## 项目模块使用

### 一、微信模块(包名:wechat)

#### 1.微信授权登录

##### 使用方法

````
package

import 	wechat_login "github.com/zhangshuai268/spg-go-pkg/service/wechat/login"

func main() {

    //初始化微信授权client
    client, err := wechat_login.NewLoginClient("appid", "appsecet")


}

````

##### 接口目录

* [完成接口文档](./doc/wechat/login.md)

#### 2.微信消息发送

##### 使用方法

````
package

import (
    "github.com/zhangshuai268/spg-go-pkg/service/wechat"
    wechat_message "github.com/zhangshuai268/spg-go-pkg/service/wechat/message"
)

func main() {
    //初始话微信消息发送client
    client, err := wechat_message.NewMessageClient("appid", "appsecet", wechat.AppTypeO) //AppTyoeO:公众号 AppTypeA:小程序
}
````

##### 接口目录

* [完成接口文档](./doc/wechat/message.md)

#### 3.微信支付

##### 使用方法

````
package

import (
    "github.com/zhangshuai268/spg-go-pkg/service/wechat"
    wechat_pay "github.com/zhangshuai268/spg-go-pkg/service/wechat/pay"
)

func main() {
    
    wechat_pay.NewPayClient("appid","mchId","paySecret") //服务商模式请区分appid和sub_appid区别
}
````

##### 接口目录

* [完成接口文档](./doc/wechat/pay.md)

### 二、阿里云模块(包名:aliyun)

#### 1.短信发送

##### 使用方法

````
package

import (
    "github.com/zhangshuai268/spg-go-pkg/service/aliyun"
    aliyun_sms "github.com/zhangshuai268/spg-go-pkg/service/aliyun/sms"
)

func main() {

    aliyun_sms.NewSmsClient("accessKeyId", "accessKeySecret", "regionId", "templateCode", "signName")
}
````

##### 接口目录

* [完成接口文档](./doc/aliyun/sms.md)
