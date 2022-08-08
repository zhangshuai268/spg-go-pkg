## 微信支付

> #### 声明：本包为gopay-sdk的封装，为方便开发封装了在开发过程中会经常出现的接口，如下单、订单查询、退款等，其他特殊功能的接口请使用gopay自行完成功能；

> #### 暂仅支持微信v2

- [gopay官方文档](https://github.com/go-pay/gopay)

> #### 微信支付相关接口，下单、订单查询、付款等，之后会继续补充

- [微信支付v2官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

### 微信支付 API

* [`client.PayUnifiedOrder()` 微信下单](#微信下单)
* [`client.PaySign()`         微信签名生成](#微信签名生成)
* [`client.PayParseNotifyToBodyMap()` 异步通知参数处理](#异步通知参数处理)
* [`client.PayVerifySign()` 验证签名](#验证签名)
* [`client.PayAddCertFilePath()` 添加退款证书](#添加退款证书)
* [`client.PayRefund()` 退款](#退款)
* [`client.PayOrderSearch()` 订单查询](#订单查询)

### 使用方法

#### 微信下单

>使用方法，参考spg-go-pkg/service/wechat/message/pay_unified_order_test.go

>文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_1.shtml

````
pm := make(spg_go_pkg.ParamMap)
pm.Set("appid", appId)
pm.Set("mch_id", mchId)
pm.Set("body", "body")
pm.Set("out_trade_no", "SASASASASAS")
pm.Set("total_fee", 0.01)
pm.Set("openid", "assas")
pm.Set("nonce_str", "sadafaf")
pm.Set("spbill_create_ip", "127.0.0.1")
pm.Set("sign_type", wechat.SignTypeMD5)
pm.Set("notify_url", "127.0.0.1/s")
pm.Set("trade_type", wechat.TradeTypeApp)
order, err := client.PayUnifiedOrder(pm)
````

#### 微信签名生成

>使用方法，参考spg-go-pkg/service/wechat/message/pay_sign_test.go

>文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=3

````
nonce := ""
packages := ""
timeStamp := ""
client.PaySign(nonce, packages, wechat.TradeTypeApp, timeStamp)
````

#### 异步通知参数处理


#### 验证签名
#### 添加退款证书
#### 退款
#### 订单查询