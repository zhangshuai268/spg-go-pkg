## 微信消息发送

>#### 微信消息发送相关接口，覆盖订阅消息、模板消息发送等方法，之后会继续补充

- [微信公众号模板消息发送官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html)
- [微信公众号订阅消息发送官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html)
- [微信小程序订阅消息发送官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html)

### 微信发送消息 API

* [`client.GetAccessToken()` 微信获取调用接口凭证(access_token)](#微信获取调用接口凭证)
* [`client.SubscribeSend()`  微信订阅消息发送(支持公众号、小程序)](#微信订阅消息发送)
* [`client.TemplateSend()`   微信模板消息发送(仅支持公众号)](#微信模板消息发送)
* [`client.TemplateId()`     微信获取模板消息id(仅支持公众号)](#微信获取模板消息id)

### 使用方法

#### 微信获取调用接口凭证

>使用方法，参考spg-go-pkg/service/wechat/message/message_token_test.go

>微信接口调用凭证的获取每日有次数限制，建议获取后储存方便下次使用，并在过期后及时更新 

````
token, err := client.GetAccessToken()
````

#### 微信订阅消息发送

>使用方法，参考spg-go-pkg/service/wechat/message/message_subscribe_send_test.go

>公众号文档地址: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html  
>小程序文档地址: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html  

````
accessTokenJs, _ := client.GetAccessToken()
openId := ""
pm := make(spg_go_pkg.ParamMap)
pm.Set("touser", openId)
pm.Set("template_id", "")
pm.Set("page", "")
pm.Set("data", map[string]interface{}{
	"name1": map[string]string{
		"value": "灰太狼",
	},
	"phrase2": map[string]string{
		"value": "通过",
	},
})
send, err := client.SubscribeSend(pm, accessTokenJs.AccessToken)
````

#### 微信模板消息发送

>使用方法，参考spg-go-pkg/service/wechat/message/message_template_send_test.go

>文档地址: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

````
accessToken, _ := client.GetAccessToken()
pm := make(spg_go_pkg.ParamMap)
pm.Set("touser", "")
pm.Set("template_id", "")
pm.Set("url", "")
pm.Set("miniprogram", map[string]interface{}{
	"appid":    "xiaochengxuappid12345",
	"pagepath": "index?foo=bar",
})
pm.Set("client_msg_id", "MSG_000001")
pm.Set("data", map[string]interface{}{
	"first": map[string]interface{}{
		"value": "恭喜你购买成功！",
		"color": "#173177",
	},
	"keyword1": map[string]interface{}{
		"value": "巧克力",
		"color": "#173177",
	},
})
send, err := client.TemplateSend(pm, accessToken.AccessToken)
````

#### 微信获取模板消息id

>使用方法，参考spg-go-pkg/service/wechat/message/message_template_id_test.go

>文档地址: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#2

````
accessToken, _ := client.GetAccessToken()
pm := make(spg_go_pkg.ParamMap)
pm.Set("template_id_short", "TM00015")
id, err := client.TemplateId(pm, accessToken.AccessToken)
````