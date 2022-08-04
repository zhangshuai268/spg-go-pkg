package aliyun

import "errors"

type SmsService interface {
	// SmsSend 短信发送
	//  支持单个或多个手机号
	//  mobile: 手机号 "1590000****或["1590000****","1350000****"]"
	//  content: 短信模板变量对应的实际值 "[{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]"
	//  文档地址: https://help.aliyun.com/document_detail/419273.html
	SmsSend(mobile string, content string) (*SmsSendResponse, error)
	// SmsSendStatic 查询短信发送统计信息
	//  isGlobal: 短信发送范围, aliyun.GlobalIn:国内短信发送记录, aliyun.GlobalOut:国际/港澳台短信发送记录
	//  startData: 开始日期，格式为yyyyMMdd，例如20181225
	//  endData: 开始日期，格式为yyyyMMdd，例如20181225
	//  pageIndex: 当前页码。默认取值为1
	//  pageSize: 每页显示的条数。取值范围：1~50
	//  文档地址: https://help.aliyun.com/document_detail/419276.html
	SmsSendStatic(isGlobal int, startData, endData string, pageIndex, pageSize int) (*SmsSendStatic, error)
	// SmsSendDetail 查询短信发送详情
	//  mobile: 手机号
	//  sendData: 短信发送日期，支持查询最近30天的记录，格式为yyyyMMdd，例如20181225
	//  pageSize: 分页查看发送记录，指定每页显示的短信记录数量，取值范围为1~50
	//  pageIndex: 分页查看发送记录，指定发送记录的当前页码
	//  文档地址: https://help.aliyun.com/document_detail/419277.html
	SmsSendDetail(mobile, sendData string, pageSize, pageIndex int) (*SmsSendDetailResponse, error)
}

type sms struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	RegionId        string `json:"region_id"`
	TemplateCode    string `json:"template_code"`
	SignName        string `json:"sign_name"`
}

// NewSmsClient 初始话短信发送客户端
//  accessKeyId: 阿里云accessKeyId
//  accessKeySecret: 阿里云accessKeySecret
//  regionId: 阿里云可用区ID
//  templateCode: 短信模板CODE
//  signName: 短信签名名称
func NewSmsClient(accessKeyId, accessKeySecret, regionId, templateCode, signName string) (SmsService, error) {
	if accessKeyId == "" {
		return nil, errors.New("缺少accessKeyId")
	}
	if accessKeySecret == "" {
		return nil, errors.New("缺少accessKeySecret")
	}
	if regionId == "" {
		return nil, errors.New("缺少regionId")
	}
	if signName == "" {
		return nil, errors.New("缺少templateCode")
	}
	if templateCode == "" {
		return nil, errors.New("缺少signName")
	}

	return &sms{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		RegionId:        regionId,
		TemplateCode:    templateCode,
		SignName:        signName,
	}, nil

}
