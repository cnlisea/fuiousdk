package fuiousdk

import (
	"bytes"
	"strconv"
)

type H5PayRequest struct {
	Version            string `json:"version"`              // 版本号
	MchntCd            string `json:"mchnt_cd"`             // 商户号
	RandomStr          string `json:"random_str"`           // 随机字符串
	OrderType          string `json:"order_type"`           // 订单类型
	OrderAmt           string `json:"order_amt"`            // 订单总金额
	MchntOrderNo       string `json:"mchnt_order_no"`       // 订单号
	TxnBeginTs         string `json:"txn_begin_ts"`         // 订单生成时间，格式为：yyyyMMddHHmmss
	GoodsDes           string `json:"goods_des"`            // 商品描述
	TermId             string `json:"term_id"`              // 终端号
	TermIp             string `json:"term_ip"`              // 终端IP
	AddnInf            string `json:"addn_inf,omitempty"`   // 附加数据
	ReservedDeviceInfo string `json:"reserved_device_info"` // 设备信息
	CurrType           string `json:"curr_type"`            // 货币类型
	Sign               string `json:"sign"`                 // 签名
	NotifyUrl          string `json:"notify_url"`           // 异步通知地址
	FrontNotifyUrl     string `json:"front_notify_url"`     // 前台通知url
	TradeType          string `json:"trade_type"`           // 交易类型 APP--app支付、H5--H5支付
}

type H5PayResponse struct {
	ResultCode        string `json:"result_code"`          // 响应代码
	ResultMsg         string `json:"result_msg"`           // 中文描述
	MchntCd           string `json:"mchnt_cd"`             // 商户代码
	TermId            string `json:"term_id"`              // 终端号
	RandomStr         string `json:"random_str"`           // 随机字符串
	OrderType         string `json:"order_type"`           // 订单类型
	QrCode            string `json:"qr_code"`              // 二维码链接
	ReservedFyOrderNo string `json:"reserved_fy_order_no"` // 富友生成的订单号
	ReservedFyTraceNo string `json:"reserved_fy_trace_no"` // 追踪号
	Sign              string `json:"sign"`                 // 签名
}

func H5Pay(cfg *Config) (*H5PayResponse, error) {
	req := &H5PayRequest{
		Version:            "1.0",
		MchntCd:            MchntCd,
		RandomStr:          RandomString(32),
		OrderType:          "WECHAT",
		OrderAmt:           strconv.FormatInt(cfg.Amount, 10),
		MchntOrderNo:       cfg.OrderNo,
		TxnBeginTs:         cfg.CreateTime.Format("20060102150405"),
		GoodsDes:           cfg.GoodsDes,
		TermId:             RandomString(8),
		TermIp:             cfg.ClientIp,
		NotifyUrl:          cfg.NotifyUrl,
		FrontNotifyUrl:     cfg.FrontNotifyUrl,
		TradeType:          "H5",
		AddnInf:            cfg.AddnInf,
		ReservedDeviceInfo: cfg.ReservedDeviceInfo,
	}

	var b bytes.Buffer
	b.WriteString(req.MchntCd)
	b.WriteString("|")
	b.WriteString(req.OrderType)
	b.WriteString("|")
	b.WriteString(req.TradeType)
	b.WriteString("|")
	b.WriteString(req.OrderAmt)
	b.WriteString("|")
	b.WriteString(req.MchntOrderNo)
	b.WriteString("|")
	b.WriteString(req.TxnBeginTs)
	b.WriteString("|")
	b.WriteString(req.GoodsDes)
	b.WriteString("|")
	b.WriteString(req.TermId)
	b.WriteString("|")
	b.WriteString(req.TermIp)
	b.WriteString("|")
	b.WriteString(req.NotifyUrl)
	b.WriteString("|")
	b.WriteString(req.RandomStr)
	b.WriteString("|")
	b.WriteString(req.Version)
	b.WriteString("|")
	b.WriteString(MchntKey)

	req.Sign = Md5(b.Bytes())

	res := new(H5PayResponse)
	if err := HttpSend(PayUrl, &req, res); err != nil {
		return nil, err
	}

	return res, nil
}
