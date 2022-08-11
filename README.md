# spg-go-pkg

## 项目介绍

此项目为spg项目组开发封装的相关工具包，为简化开发时相关处理代码，因此封装此包，它将会经常更新。

> #### 本包中所有方法均有参数注释，在调用时请注意相关代码注释或者代码编译工具的提示！

> #### >使用时每引入一个新包，执行go mod tidy，来下载包内的相关依赖

## 项目下载

````
go get github.com/zhangshuai268/spg-go-pkg
````

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

#### 2.支付宝支付

##### 使用方法


````
package

import (
    "github.com/zhangshuai268/spg-go-pkg/service/aliyun"
    aliyun_pay "github.com/zhangshuai268/spg-go-pkg/service/aliyun/pay"
)

func main() {
    
    //aliyun.TradeTypeApp: APP支付
    //aliyun.TradeTypeWeb: 手机网站支付
    //aliyun.TradeTypeH5:  电脑网站支付
    aliyun_sms.NewPayClient("appId", "privateKey", "publicKey", aliyun.TradeTypeApp, isProd)
}
````

##### 接口目录

* [完成接口文档](./doc/aliyun/pay.md)

### 三、工具包

#### 1.logger包

> 封装包: github.com/sirupsen/logrus

> 官方文档: https://github.com/sirupsen/logrus

> main中初始化后便可全局调用

> 注意: 若文件记录开关打开会在项目根目录下生成log文件夹，日记记录在log文件夹中

##### 使用方法

````
import "github.com/zhangshuai268/spg-go-pkg/pkg/logger"

func main() {

    //初始化日志
	_, err := logger.InitLogger(true)
	if err != nil {
		panic("日志初始化失败" + err.Error())
	}
	
}
````

##### 封装函数

* `logger.Logger.Info(...interface{})` 自定义日志记录,通常用于debug  
* `logger.Logger.Error(...interface{})` 错误类型日志记录,通常用于记录error类型错误  
* `logger.Logger.Panic(...interface{})` 异常类型日志记录,通常用于全局检测异常  
* `logger.Logger.Api(...interface{})` 接口调用日志记录,通常用于记录参数、ip等接口调用信息  

#### 2.iarray包

> 封装了相关int[]处理函数

##### 封装函数

* `iarray.IsContain([]int, int) (bool, int) ` 检查int是否在[]int数组中
* `iarray.Reverse(*int[])` 反转数组
* `iarray.RemoveReplica([]int) []int` 数组去重
* `iarray.Union([]int, []int) []int` 求并集
* `iarray.Intersect([]int, []int) []int` 求交集

#### 3.format包

> 封装了相关号码判断函数

##### 封装函数

* `format.VerifyMobileFormat(string) bool` 手机号格式判断
* `format.VerifyEmailFormat(string) bool` 邮箱格式判断
* `format.VerifyIdFormat(string) bool` 身份证号格式判断

#### 4.myTime包

> 自定义时间类型， myTime.MyTime

> time.Time类型在输出时，默认格式为2022-08-10T11:22:08+08:00，数据处理较为麻烦，因此自定义时间，myTime.MyTime默认格式为2006-01-02 15:04:05

> myTime.MyTime可与time.Time类型相互进行强制类型转换

##### 封装函数

* `myTime.StringToMyTime(string)` 字符串转myTime.MyTime类型

#### 5.util包

> 工具包，封装了一些常用的函数，之后与其他包无关的函数均会封装在此包内，可推荐一些好用的函数进行封装

##### 封装函数

* `util.StructTo(old interface{}, new interface{}) error` 类型转换
* `util.GetRandNum(n int, char ...string) string` 生成随机字符串
* `util.Md5Make(s string) string` MD5加密

#### 6.xorm包

> 封装包: github.com/go-xorm/xorm

> 官方文档: https://gitea.com/xorm/xorm/src/branch/master/README_CN.md

> 主要封装了生成xorm的调用engine，之后的操作，请参考官方文档

##### 使用方法

````
import "github.com/zhangshuai268/spg-go-pkg/pkg/xorm"

func main() {

  //xormEngine.Orm为xorm的调用engine
  xormEngine, err := xorm.InitXormEngine("mysql", "user:password@(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=true&loc=Local&collation=utf8mb4_unicode_ci")

}
````

#### 7.gorm包

> 封装包: gorm.io/gorm

> 文档: https://learnku.com/docs/gorm/v2/index/9728

> 主要封装了生成gorm的调用engine，之后的操作，请参考官方文档

##### 使用方法

````
import "github.com/zhangshuai268/spg-go-pkg/pkg/gorm"

func main() {

  //gormEngine.Orm为xorm的调用engine
  gormEngine, err := gorm.InitGormEngine("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")

}
````

#### 8.redis包

> 封装包： github.com/go-redis/redis/v8

> 仅支持Redis 6及一下版本

> 官方文档: https://redis.uptrace.dev/

> 主要封装了生成redis的调用engine，之后的操作，请参考官方文档

##### 使用方法

````
import "github.com/zhangshuai268/spg-go-pkg/pkg/redis"

func main() {

  //redisEngine.Orm为redis的调用engine
  redisEngine, err := redis.InitRedisEngine("127.0.0.1:6397", "", 0)
  
}
````

#### 9.mongodb包


> 封装包： go.mongodb.org/mongo-driver

> 官方文档: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo

> 主要封装了生成mongo的调用engine，之后的操作，请参考官方文档

````
import "github.com/zhangshuai268/spg-go-pkg/pkg/mongo"

func main() {

  //mongoEngin.Orm为mongo的调用engine
  mongoEngin, err := mongo.InitMongoEngin("mongodb://User:PassWord@127.0.0.1:27017", "dbname")
  
}
````

#### 10.http包

> 封装了相关http请求方法

##### 封装函数

* `http.HttpGet(url string) (map[string]interface{}, error)` get请求
* `http.HttpPostJson(url string, post []byte) (map[string]interface{}, error)` post请求json格式
* `http.HttpPostForm(urls string, post map[string]string) (map[string]interface{}, error)` post请求form格式
* `http.GetIP(r *http.Request) (string, error)` 获取http请求IP



