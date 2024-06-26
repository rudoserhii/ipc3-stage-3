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
	return &Models{}
}

func (p *Models) Product() interfaces.ProductModel {
	return p.product
}

func (o *Models) Order() interfaces.OrderModel {
	return o.order
}
