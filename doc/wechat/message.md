## 微信消息发送

>#### 微信消息发送相关接口，覆盖订阅消息、模板消息发送等方法，之后会继续补充

- [微信公众号模板消息发送官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html)
- [微信公众号订阅消息发送官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html)
- [微信小程序订阅消息发送官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html)

### 微信发送消息 API

* `client.GetAccessToken()` 微信获取调用接口凭证(access_token)
* `client.SubscribeSend()`  微信订阅消息发送(支持公众号、小程序)
* `client.TemplateSend()`   微信模板消息发送(仅支持公众号)
* `client.TemplateId()`     微信获取模板消息id(仅支持公众号)