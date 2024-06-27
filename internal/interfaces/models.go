package interfaces

import "gorm.io/gorm"

// Interface Models aggregates puts together all Models.
type Models interface {
	Product() ProductModel
	Order() OrderModel
}

// Type Product is a Model
type Product struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Image       string  `gorm:"not null"`
	Available   bool    `gorm:"not null"`
	Orders      []Order `gorm:"foreignKey:ProductID"`
}

// Type Order is a Model
type Order struct {
	gorm.Model
	ProductID uint
	Quantity  uint `gorm:"not null"`
}

type ProductModel interface {
	GetProductByID(productID uint) (*Product, error)
	GetAllProducts() ([]Product, error)
	CreateProduct(p *Product) error
	UpdateProductByID(p *Product, productID uint) error
	DeleteProductByID(productID uint) error
}

type OrderModel interface {
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID uint) (*Order, error)
	CreateOrder(order *Order) error
	CancelOrderByID(orderID uint) error
}
