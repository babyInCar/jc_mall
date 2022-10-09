package model

type GoodsDetail struct {
	Goods int32
	Num   int32
}

// type GoodsDetailList []GoodsDetail
type GoodsDetailList []GoodsDetail

type Invetory struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stock   int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int";comment:'类似于分布式锁的乐观锁'`
}

type InventoryNew struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"` //分布式锁的乐观锁
	Freeze  int32 `gorm:"type:int"` //冻结库存
}

type Delivery struct {
	Goods   int32  `gorm:"type:int;index"`
	Nums    int32  `gorm:"type:int"`
	OrderSn string `gorm:"type:varchar(200)"`
	Status  string `gorm:"type:varchar(200);comment:'状态：1.等待支付 2.支付成功 3.失败'"`
}

type StockSellDetail struct {
	// OrderSn string `gorm:"type:varchar(200);index:idx_order_sn,unique;"`
	// Status int32 `gorm:"type:varchar(200)"` //1 表示已扣减 2. 表示已归还
	// Detail GoodsDetailList `gorm:"type:varchar(200)"`
	OrderSn string          `gorm:"type:varchar(200);index:index_order_sn,unique;"`
	Status  int32           `gorm:"type:varchar(200);comment:'1：已扣减 2.表示已归还'"`
	Detail  GoodsDetailList `gorm:"type:varchar(200)"`
}

func (StockSellDetail) TableName() string {
	return "stockselldetail"
}
