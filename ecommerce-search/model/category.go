package model

import (
	"time"
)

type Category struct {
	CategoryID       uint      `gorm:"primaryKey;column:category_id" json:"category_id"`
	CategoryName     string    `gorm:"column:category_name" json:"category_name"`
	ParentCategoryID *uint     `gorm:"column:parent_category_id" json:"parent_category_id,omitempty"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	Products         []Product `gorm:"many2many:product_categories;" json:"products,omitempty"`
}
