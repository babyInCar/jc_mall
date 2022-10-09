package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	id         int32          `gorm:"primarykey;type:int" json:"id"`
	CreateTime time.Time      `gorm:"column:add_time" json:"-"`
	UpdateTime time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	IsDeleted  bool           `json:"-"`
}
