package wechat

import "strings"

const (
	LandCN   = "zh_CN"    //简体
	LandTW   = "zh_TW"    //繁体
	LandEN   = "en"       //英语
	AppTypeA = "Applets"  //小程序
	AppTypeO = "Official" //公众号

	// 支付类型
	TradeTypeMini   = "JSAPI"  // 小程序支付
	TradeTypeJsApi  = "JSAPI"  // JSAPI支付
	TradeTypeApp    = "APP"    // app支付
	TradeTypeH5     = "MWEB"   // H5支付
	TradeTypeNative = "NATIVE" // Native支付

	// 签名方式
	SignTypeMD5        = "MD5"
	SignTypeHMACSHA256 = "HMAC-SHA256"

	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

type NotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// ToXmlString 返回数据给微信
func (w *NotifyResponse) ToXmlString() (xmlStr string) {
	var buffer strings.Builder
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(w.ReturnCode)
	buffer.WriteString("]]></return_code>")
	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(w.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}
