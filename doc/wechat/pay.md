## 微信支付

> #### 声明：本包为gopay-sdk的封装，为方便开发封装了在开发过程中会经常出现的接口，如下单、订单查询、退款等，其他特殊功能的接口请使用gopay自行完成功能；

> #### 暂仅支持微信v2

- [gopay官方文档](https://github.com/go-pay/gopay)

> #### 微信支付相关接口，下单、订单查询、付款等，之后会继续补充

- [微信支付v2官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

### 微信支付 API

* [`client.PayUnifiedOrder()` 微信下单](#微信下单)
* [`client.PaySign()`         微信签名生成](#微信签名生成)
* [`client.PayParseNotifyToBodyMap()` 微信异步通知参数处理](#微信异步通知参数处理)
* [`client.PayVerifySign()` 微信验证签名](#微信验证签名)
* [`client.PayAddCertFilePath()` 微信添加退款证书](#微信添加退款证书)
* [`client.PayRefund()` 微信退款](#微信退款)
* [`client.PayOrderSearch()` 微信订单查询](#微信订单查询)

### 使用方法

#### 微信下单

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_unified_order_test.go

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

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_sign_test.go

>文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=3

````
nonce := ""
packages := ""
timeStamp := ""
client.PaySign(nonce, packages, wechat.TradeTypeApp, timeStamp)
````

#### 微信异步通知参数处理

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_notify_bodymap_test.go

````
//ctx: *gin.Context
notifyReq, err := client.PayParseNotifyToBodyMap(ctx.Request)
````

#### 微信验证签名

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_verify_sign_test.go

````
//ctx: *gin.Context
notifyReq, err := client.PayParseNotifyToBodyMap(ctx.Request)
sign, err := client.PayVerifySign(wechat.SignTypeMD5, notifyReq)
````

#### 微信添加退款证书

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_add_cert_file_test.go

````
certFilePath := ""
keyFilePath := ""
err := client.PayAddCertFilePath(certFilePath, keyFilePath)
````

#### 微信退款

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_refund_test.go

>文档地址: https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_4.shtml

>注意: 微信支付调用退款需要添加退款证书

````
err := client.PayAddCertFilePath("", "")
if err != nil {
	return
}
pm := make(spg_go_pkg.ParamMap)
pm.Set("appid", appId)
pm.Set("mch_id", mchId)
pm.Set("out_trade_no", "SASASASASAS")
pm.Set("out_refund_no", "sdaa")
pm.Set("total_fee", 1)
pm.Set("refund_fee", 1)
refund, resMap, err := client.PayRefund(pm)
````

#### 微信订单查询

>使用方法，参考spg-go-pkg/service/wechat/pay/pay_refund_test.go

>文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_2.shtml

````
pm := make(spg_go_pkg.ParamMap)
pm.Set("appid", appId)
pm.Set("mch_id", mchId)
pm.Set("transaction_id", "SASASASASAS")
pm.Set("nonce_str", "sadafaf")
search, resMap, err := client.PayOrderSearch(pm)
````

### 其他

#### 回调接口处理方法

>使用gin框架

````
func Callback(ctx *gin.Context) {
//回复微信的数据
rsp := new(spg_go_pkg.NotifyResponse)

}
````