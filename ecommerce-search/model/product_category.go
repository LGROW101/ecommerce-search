package model

import (
	"time"
)

type ProductCategory struct {
	ProductID  uint      `gorm:"primaryKey;column:product_id" json:"product_id"`
	CategoryID uint      `gorm:"primaryKey;column:category_id" json:"category_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}
