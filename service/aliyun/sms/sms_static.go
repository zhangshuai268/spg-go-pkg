package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
)

type SmsSendStatic dysmsapi.QuerySendStatisticsResponse

func (s *sms) SmsSendStatic(isGlobal int, startData, endData string, pageIndex, pageSize int) (*SmsSendStatic, error) {
	var res SmsSendStatic
	client, err := dysmsapi.NewClientWithAccessKey(s.RegionId, s.AccessKeyId, s.AccessKeySecret)
	request := dysmsapi.CreateQuerySendStatisticsRequest()
	request.Scheme = "https"
	request.IsGlobe = requests.NewInteger(isGlobal)
	request.StartDate = startData
	request.EndDate = endData
	request.PageIndex = requests.NewInteger(pageIndex)
	request.PageSize = requests.NewInteger(pageSize)
	statistics, err := client.QuerySendStatistics(request)
	if err != nil {
		return nil, err
	}
	err = util.StructTo(statistics, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil

}
