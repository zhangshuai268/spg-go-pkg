## 微信登录

>#### 微信授权登录相关接口，覆盖网页授权、小程序授权、手机号解密等方法，之后会继续补充

- [微信网页授权官方文档](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)
- [微信小程序授权官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/login.html)
- [手机号解密算法官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)

### 微信登录 API

* `client.Oauth(code string)`  网页授权登录
* `client.UserInfo(accessToken, openId, lang string)` 获取用户信息
* `client.JsCode2(code string)` 小程序授权登录
* `client.Auth(access_token, open_id string)` access_token有效期判断
* `client.Mobile(mobile, iv, sessionKey string)` 手机号解密