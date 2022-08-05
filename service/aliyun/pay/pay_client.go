package aliyun

import (
	"errors"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"net/http"
)

type PayService interface {
	// PayUnifiedOrder 下单
	//  App支付官方文档地址：https://opendocs.alipay.com/open/02e7gq?scene=20
	//  手机网站支付官方文档地址：https://opendocs.alipay.com/open/02e7gq?scene=21
	//  电脑网站支付官方文档地址：https://opendocs.alipay.com/open/02e7gq?scene=22
	PayUnifiedOrder(pm spg_go_pkg.ParamMap) (string, error)
	// PayParseNotifyToBodyMap 异步通知参数处理
	//  解析支付宝支付异步通知的结果参数
	PayParseNotifyToBodyMap(req *http.Request) (spg_go_pkg.ParamMap, error)
	// PayVerifySign 校验签名
	//  用于同步或异步验证签名
	//  bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq，推荐通 ParamMap 验签
	//  返回参数ok：是否验签通过
	PayVerifySign(bean interface{}) (bool, error)
	// PayRefund 退款接口
	//  文档地址：https://opendocs.alipay.com/open/02e7go
	PayRefund(pm spg_go_pkg.ParamMap) (*PayRefundResponse, error)
	// PayOrderSearch 订单查询接口
	//  文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_2.shtml
	PayOrderSearch(pm spg_go_pkg.ParamMap) (*PayOrderResponse, error)
}

type pay struct {
	AppId      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	IsProd     bool   `json:"is_prod"`
	PayType    string `json:"pay_type"`
}

// NewPayClient 初始化支付客户端
//  appId: 支付宝分配给开发者的应用ID
//  privateKey: 应用私钥
//  publicKey:  应用公钥
//  payType:  支付类型，现只支持，App: aliyun.TradeTypeApp 手机网站: aliyun.TradeTypeWeb 电脑网站: aliyun.TradeTypeH5
//  isProd: 环境，true:正式环境 false: 测试环境
func NewPayClient(appId, privateKey, publicKey, payType string, isProd bool) (PayService, error) {
	if appId == "" {
		return nil, errors.New("缺少appid")
	}
	if privateKey == "" {
		return nil, errors.New("缺少privateKey")
	}
	if publicKey == "" {
		return nil, errors.New("缺少publicKey")
	}
	return &pay{
		AppId:      appId,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		IsProd:     isProd,
		PayType:    payType,
	}, nil

}
