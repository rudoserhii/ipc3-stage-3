package models

import (
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"gorm.io/gorm"
)

type OrderModel struct {
	DB *gorm.DB
}

func (o *OrderModel) GetAllOrders() ([]interfaces.Order, error) {
	var orders []interfaces.Order
	if err := o.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderModel) GetOrderByID(orderID uint) (*interfaces.Order, error) {
	var order interfaces.Order
	if err := o.DB.First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderModel) CreateOrder(order *interfaces.Order) error {
	if err := o.DB.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (o *OrderModel) CancelOrderByID(orderID uint) error {
	if err := o.DB.Delete(&interfaces.Order{}, orderID).Error; err != nil {
		return err
	}
	return nil
}
