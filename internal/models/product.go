package models

import "github.com/obiMadu/ipc3-stage-3/internal/interfaces"

type ProductModel struct{}

func (p *ProductModel) GetAllProducts() ([]interfaces.Product, error) {
	return []interfaces.Product{}, nil
}

func (p *ProductModel) GetProductByID(productID uint) (*interfaces.Product, error) {
	return &interfaces.Product{}, nil
}

func (p *ProductModel) CreateProduct(product *interfaces.Product) error {
	return nil
}

func (p *ProductModel) UpdateProductByID(product *interfaces.Product) error {
	return nil
}

func (p *ProductModel) DeleteProductByID(productID uint) error {
	return nil
}
