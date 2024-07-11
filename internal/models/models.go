package models

import (
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"gorm.io/gorm"
)

type Models struct {
	product interfaces.ProductModel
	order   interfaces.OrderModel
}

func (m *Models) Product() interfaces.ProductModel {
	return m.product
}

func (m *Models) Order() interfaces.OrderModel {
	return m.order
}

func NewModels(db *gorm.DB) interfaces.Models {
	return &Models{
		product: &ProductModel{DB: db},
		order:   &OrderModel{DB: db},
	}
}
