package model

import (
	"time"
)

type Product struct {
	ProductID     uint       `gorm:"primaryKey;column:product_id" json:"product_id"`
	ProductName   string     `gorm:"column:product_name" json:"product_name"`
	Description   string     `gorm:"column:description" json:"description"`
	Price         float64    `gorm:"column:price" json:"price"`
	StockQuantity int        `gorm:"column:stock_quantity" json:"stock_quantity"`
	SearchVector  string     `gorm:"column:search_vector" json:"-"`
	SearchText    string     `gorm:"column:search_text" json:"-"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
	Categories    []Category `gorm:"many2many:product_categories;joinForeignKey:product_id;joinReferences:category_id" json:"categories,omitempty"`
}
