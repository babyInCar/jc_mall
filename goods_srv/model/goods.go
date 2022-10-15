package model

type Category struct {
	BaseModel
	// Name  string `gorm:"type:varchar(20);not null" json:"name"`
	// ParentCategoryID int32 `json:"parent"`
	// ParentCategory *Category `json:"-"`
	// SubCategory []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	// Level int32 `gorm:"type:int;not null;default:1" json:"level"`
	// IsTab bool `gorm:"default:false;not null" json:"is_tab"`
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignkey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category
	BrandID    int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
	// Name string `gorm:"type:varchar(20)"`
}

type GoodsModel struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:default:false;not null`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:default:false;not null`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodsSn         string   `gorm:"type:varchar(50) comment '商品编码';not nul"`
	ClickNum        int32    `gorm:"type:int comment '点击数量';default:0;not null"`
	SoldNum         int32    `gorm:type:int;default:0;not null`
	FavNum          int32    `gorm:"type:int comment:'收藏数量';default 0;not null"`
	MarketPrice     float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsDesc       string   `gorm:"type:varchar(100);not null"`
	Images          GormList `gorm:"type:varhcar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsCoverImage string   `gorm:"type:varchar(200);not null"`
}
