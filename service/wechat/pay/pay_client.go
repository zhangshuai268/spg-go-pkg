package wechat

import (
	"errors"
	spg_go_pkg "github.com/zhangshuai268/spg-go-pkg"
	"net/http"
)

type PayService interface {
	// PayUnifiedOrder 下单
	//  文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_1.shtml
	PayUnifiedOrder(pm spg_go_pkg.ParamMap) (*PayUnifiedOrderResponse, error)
	// PaySign 签名生成
	//  JSAPI，用于统一下单获取支付参数后，再次计算出小程序用的paySign
	//  nonceStr：随即字符串
	//  packages：统一下单成功后拼接得到的值
	//  signType：签名类型
	//  timeStamp：时间
	//  微信小程序支付PaySign计算文档：https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=3
	PaySign(nonceStr string, packages string, signType string, timeStamp string) string
	// PayParseNotifyToBodyMap 异步通知参数处理
	//  解析微信支付异步通知的结果参数
	PayParseNotifyToBodyMap(req *http.Request) (spg_go_pkg.ParamMap, error)
	// PayVerifySign 校验签名
	//  用于同步或异步验证签名
	//  signType：签名类型（调用API方法时填写的类型）
	//  bean：微信同步返回的结构体 wxRsp 或 异步通知解析的结构体 notifyReq，推荐通 ParamMap 验签
	//  返回参数ok：是否验签通过
	PayVerifySign(signType string, bean interface{}) (bool, error)
	// PayAddCertFilePath 添加退款证书
	//  添加退款证书
	PayAddCertFilePath(certFilePath string, keyFilePath string) error
	// PayRefund 退款接口
	//  注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
	//  文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_4.shtml
	PayRefund(pm spg_go_pkg.ParamMap) (*PayRefundResponse, spg_go_pkg.ParamMap, error)
	// PayOrderSearch 订单查询接口
	//  文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_2.shtml
	PayOrderSearch(pm spg_go_pkg.ParamMap) (*PayOrderResponse, spg_go_pkg.ParamMap, error)
}

type pay struct {
	AppId        string `json:"app_id"`
	MchId        string `json:"mch_id"`
	PaySecret    string `json:"pay_secret"`
	CertFilePath string `json:"cert_file_path"`
	KeyFilePath  string `json:"key_file_path"`
}

// NewPayClient 初始化支付客户端
//  appId: 微信支付分配的公众账号ID,服务商模式根据文档进行传递
//  mchId: 商户id
//  paySecret: 支付密钥
func NewPayClient(appId, mchId, paySecret string) (PayService, error) {
	if appId == "" {
		return nil, errors.New("缺少appid")
	}
	if mchId == "" {
		return nil, errors.New("缺少mch_id")
	}
	if paySecret == "" {
		return nil, errors.New("缺少pay_secret")
	}

	return &pay{
		AppId:     appId,
		MchId:     mchId,
		PaySecret: paySecret,
	}, nil

}
