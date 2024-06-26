package models

import "github.com/obiMadu/ipc3-stage-3/internal/interfaces"

type OrderModel struct{}

func (o *OrderModel) GetAllOrders() ([]interfaces.Order, error) {
	return []interfaces.Order{}, nil
}

func (o *OrderModel) GetOrderByID(orderID uint) (interfaces.Order, error) {
	return interfaces.Order{}, nil
}

func (o *OrderModel) CreateOrder(order *interfaces.Order) error {
	return nil
}

func (o *OrderModel) CancelOrderByID(orderID uint) error {
	return nil
}
