package Model

import "time"

type ShoppingCart struct {
	BaseModel
}

type OrderInfo struct {
	BaseModel

	User       int32  `gorm:"type:int;index"`
	OrderSn    string `gorm:"type:varhcar(30) comment:'订单编号';index"`
	PayType    string `gorm:"type:"varchar(20) comment:'alipay, wechat'"`
	Status     string `gorm:"type:varchar(20) comment:'PAYING, TRADE_SUCCESS,TRADE_CLOSED,WAIT_BUYER_PAY,TRADE_FINISHED'"`
	TradeNo    string `gorm:"type:varchar(100) comment '交易流水号'"`
	OrderMount float32
	PayTime    *time.Time `gorm:"type:datetime comment '支付时间'"`

	Address      string `gorm:"type:varchar(100)"`
	SignerName   string `gorm:"type:varchar(20)"`
	SingerMobile string `gorm:type:varchar(11)`
	Post         string `gorm:"type:varchar(20)"`
}

func (OrderInfo) TableName() string {
	return "orderinfo"
}

type OrderGoods struct {
	BaseModel

	Order int32 `gorm:"type:int;index"`
	Goods int32 `gorm:"type:int;index"`

	//商品的信息保存下来，字段冗余，高并发系统中一般都不会遵循三范式 作镜像 记录
	GoodsName  string `gorm:"type:varchar(100) comment:'商品名称;index"`
	GoodsImage string `gorm:"type:varchar(200)"`
	GoodsPrice float32
	Nums       int32 `gorm:"type:int"`
}

func (OrderGoods) TableName() string {
	return "ordergoods"
}
