package fuiousdk

import (
	"testing"
	"time"
)

func TestH5Pay(t *testing.T) {
	cfg := &Config{
		OrderNo: "106620180514"+"111111111111111119",
		Amount: 1,
		CreateTime: time.Now(),
		GoodsDes: "元来房卡测试",
		ClientIp: "127.0.0.1",
		NotifyUrl: "http://test.yuanlaihuyu.com/pay/notify",
		ReservedDeviceInfo: `{"type":"IOS","app_name":"元来棋牌","app_url":"com.yuanlai.majiangtest"}`,
	}
	res, err := H5Pay(cfg)
	if err != nil {
		t.Fatal("h5 pay fail", err)
	}

	t.Log("res:", res)
}
