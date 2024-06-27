package models

import (
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Models struct {
	product *ProductModel
	order   *OrderModel
}

func NewModels(db *gorm.DB) interfaces.Models {
	DB = db
	return &Models{
		product: &ProductModel{},
		order:   &OrderModel{},
	}
}

func (m *Models) Product() interfaces.ProductModel {
	return m.product
}

func (m *Models) Order() interfaces.OrderModel {
	return m.order
}
