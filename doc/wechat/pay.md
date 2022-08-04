## 微信支付

> #### 声明：本包为gopay-sdk的封装，为方便开发封装了在开发过程中会经常出现的接口，如下单、订单查询、退款等，其他特殊功能的接口请使用gopay自行完成功能；

> #### 暂仅支持微信v2

- [gopay官方文档](https://github.com/go-pay/gopay)

> #### 微信支付相关接口，下单、订单查询、付款等，之后会继续补充

- [微信支付v2官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)

### 微信支付 API

* `client.PayUnifiedOrder()` 微信下单
* `client.PaySign()`         微信签名生成
* `client.PayParseNotifyToBodyMap()` 异步通知参数处理
* `client.PayVerifySign()` 验证签名
* `client.PayAddCertFilePath()` 添加退款证书
* `client.PayRefund()` 退款
* `client.PayOrderSearch()` 订单查询