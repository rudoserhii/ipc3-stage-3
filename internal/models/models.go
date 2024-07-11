package models

import (
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"gorm.io/gorm"
)

type Models struct {
	productModel interfaces.ProductModel
	orderModel   interfaces.OrderModel
}

func (m *Models) Product() interfaces.ProductModel {
	return m.productModel
}

func (m *Models) Order() interfaces.OrderModel {
	return m.orderModel
}

func NewModels(db *gorm.DB) interfaces.Models {
	return &Models{
		productModel: &ProductModel{DB: db},
		orderModel:   &OrderModel{DB: db},
	}
}
