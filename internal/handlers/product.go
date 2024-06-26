package handlers

import "github.com/gin-gonic/gin"

type ProductRoutes struct{}

func (p *ProductRoutes) GetAllProducts(*gin.Context) {
}

func (p *ProductRoutes) GetProductByID(*gin.Context) {
}

func (p *ProductRoutes) CreateProduct(*gin.Context) {
}

func (p *ProductRoutes) UpdateProductByID(*gin.Context) {
}

func (p *ProductRoutes) DeleteProductByID(*gin.Context) {
}
