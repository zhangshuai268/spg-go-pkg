## 短信发送

> #### 声明: 本包为阿里云go-sdk的封装，为方便开发封装了在开发过程中会经常出现的接口，如短信发送、短信查询、短信统计查询等，其他特殊功能的接口请使用阿里云go-sdk自行完成功能；

- [阿里云go-sdk官方文档](https://help.aliyun.com/document_detail/311619.html)

> #### 短信发送、短信查询、短信统计查询等，之后会继续补充

- [阿里云短信发送官方文档](https://help.aliyun.com/document_detail/419273.html)

### 短信发送 API

* [`client.SmsSend()` 短信发送](#短信发送)
* [`client.SmsSendStatic()` 查询短信发送统计信息](#查询短信发送统计信息)
* [`client.SmsSendDetail()` 查询短信发送详情](#查询短信发送详情)

### 使用方法

#### 短信发送

>使用方法，参考spg-go-pkg/service/wechat/sms/sms_send_test.go

>文档地址: https://help.aliyun.com/document_detail/419273.html

````
content := `{"code":"999999"}`
send, err := client.SmsSend("173****1241", content)
````

#### 查询短信发送统计信息

>使用方法，参考spg-go-pkg/service/wechat/sms/sms_static_test.go

>文档地址: https://help.aliyun.com/document_detail/419276.html

````
//GlobalIn:国内 GlobalOut:国外或港澳台
static, err := client.SmsSendStatic(aliyun.GlobalIn, "20220804", "20220804", 1, 10)
````

#### 查询短信发送详情

>使用方法，参考spg-go-pkg/service/wechat/sms/sms_detail_test.go

>文档地址: https://help.aliyun.com/document_detail/419277.html

````
res, err := client.SmsSendDetail("123****2313", "20200801", 10, 1)
````