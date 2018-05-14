package fuiousdk

import "time"

type Config struct {
	OrderNo            string    // 订单号
	Amount             int64     // 金额, 单位分
	CreateTime         time.Time // 订单生成时间
	GoodsDes           string    // 商品简要描述
	ClientIp           string    // 终端ip
	NotifyUrl          string    // 异步通知地址
	FrontNotifyUrl     string    // 前台通知url
	AddnInf            string    // 附加数据
	ReservedDeviceInfo string    // 设备信息
}
