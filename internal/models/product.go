package models

import (
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"gorm.io/gorm"
)

type ProductModel struct {
	DB *gorm.DB
}

func (p *ProductModel) GetAllProducts() ([]interfaces.Product, error) {
	var products []interfaces.Product
	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductModel) GetProductByID(productID uint) (*interfaces.Product, error) {
	var product interfaces.Product
	if err := p.DB.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductModel) CreateProduct(product *interfaces.Product) error {
	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductModel) UpdateProductByID(product *interfaces.Product, productID uint) error {
	if err := p.DB.Model(&interfaces.Product{}).Where("id = ?", productID).Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductModel) DeleteProductByID(productID uint) error {
	if err := p.DB.Delete(&interfaces.Product{}, productID).Error; err != nil {
		return err
	}
	return nil
}
