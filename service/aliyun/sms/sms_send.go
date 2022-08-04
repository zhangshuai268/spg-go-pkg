package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type SmsSendResponse dysmsapi.SendSmsResponse

func (s *sms) SmsSend(mobile, content string) (*SmsSendResponse, error) {
	var res SmsSendResponse
	client, err := dysmsapi.NewClientWithAccessKey(s.RegionId, s.AccessKeyId, s.AccessKeySecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile         //接收短信的手机号码
	request.SignName = s.SignName         //短信签名名称
	request.TemplateCode = s.TemplateCode //短信模板ID
	request.TemplateParam = content
	response, err := client.SendSms(request)
	if err != nil {
		return nil, err
	}
	err = util.StructTo(response, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
