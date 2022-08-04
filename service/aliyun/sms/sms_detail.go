package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type SmsSendDetailResponse dysmsapi.QuerySendDetailsResponse

func (s *sms) SmsSendDetail(mobile, sendData string, pageSize, pageIndex int) (*SmsSendDetailResponse, error) {
	var res SmsSendDetailResponse
	client, err := dysmsapi.NewClientWithAccessKey(s.RegionId, s.AccessKeyId, s.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	request := dysmsapi.CreateQuerySendDetailsRequest()
	request.Scheme = "https"
	request.PhoneNumber = mobile
	request.SendDate = sendData
	request.PageSize = requests.NewInteger(pageSize)
	request.CurrentPage = requests.NewInteger(pageIndex)
	details, err := client.QuerySendDetails(request)
	if err != nil {
		return nil, err
	}
	err = util.StructTo(details, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
