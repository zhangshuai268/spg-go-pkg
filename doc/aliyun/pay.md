## 支付宝支付

> #### 声明：本包为gopay-sdk的封装，为方便开发封装了在开发过程中会经常出现的接口，如下单、订单查询、退款等，其他特殊功能的接口请使用gopay自行完成功能；

- [gopay官方文档](https://github.com/go-pay/gopay)

> #### 支付宝支付接口种类较多，现已完成: 手机网站支付、电脑网站支付、手机APP支付

> #### 支付宝支付相关接口，下单、订单查询、付款等，之后会继续补充

- [支付宝支付官当文档](https://opendocs.alipay.com/open/00a0ut)

### 支付宝支付 API

* [`client.PayUnifiedOrder()` 支付宝下单](#支付宝下单)
* [`client.PayParseNotifyToBodyMap()` 支付宝异步通知参数处理](#支付宝异步通知参数处理)
* [`client.PayVerifySign()` 支付宝校验签名](#支付宝校验签名)
* [`client.PayRefund()` 支付宝退款](#支付宝退款)
* [`client.PayOrderSearch()` 支付宝订单查询](#支付宝订单查询)

### 使用方法

#### 支付宝下单

>使用方法，参考spg-go-pkg/service/aliyun/pay/pay_unified_order_test.go

>App支付官方文档地址：https://opendocs.alipay.com/open/02e7gq?scene=20  
>手机网站支付官方文档地址：https://opendocs.alipay.com/open/02e7gq?scene=21  
>电脑网站支付官方文档地址：https://opendocs.alipay.com/open/02e7gq?scene=22  

````
pm := make(spg_go_pkg.ParamMap)
pm.Set("out_trade_no", "sadasdadasd")
pm.Set("total_amount", 0.01)
pm.Set("subject", "测试")
pm.Set("product_code", "QUICK_WAP_WAY")
pm.Set("return_url", "www.baidu.com")
pm.Set("notify_url", "www.baidu.com")
order, err := client.PayUnifiedOrder(pm)
````

#### 支付宝异步通知参数处理

>使用方法，参考spg-go-pkg/service/aliyun/pay/pay_notify_bodymap_test.go

````
//ctx: *gin.Context
notifyReq, err := client.PayParseNotifyToBodyMap(ctx.Request)
````

#### 支付宝校验签名

>使用方法，参考spg-go-pkg/service/aliyun/pay/pay_verify_sign_test.go

````
//ctx: *gin.Context
notifyReq, err := client.PayParseNotifyToBodyMap(ctx.Request)
sign, err := client.PayVerifySign(notifyReq)
````

#### 支付宝退款

>使用方法，参考spg-go-pkg/service/aliyun/pay/pay_refund_test.go

>文档地址: https://opendocs.alipay.com/open/02e7go

````
pm := make(spg_go_pkg.ParamMap)
pm.Set("out_trade_no", "sadasdadasd")
pm.Set("refund_amount", 0.01)
order, err := client.PayRefund(pm)
````

#### 支付宝订单查询

>使用方法，参考spg-go-pkg/service/aliyun/pay/pay_order_search_test.go

>文档地址: https://opendocs.alipay.com/open/02e7gm

````
pm := make(spg_go_pkg.ParamMap)
pm.Set("out_trade_no", "sadasdadasd")
search, err := client.PayOrderSearch(pm)
````

### 其他

#### 回调接口处理方法
