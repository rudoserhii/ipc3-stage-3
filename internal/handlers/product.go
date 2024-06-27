package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
)

var uploadDir string = "./images/products/"
var imgRoute string = "/products/images/"

type ProductRoutes struct{}

// @Summary			Get all products in store
// @Description		Get all products in store
// @Tags			Products
// @ID				get-all-products
// @Produce			json
// @Failure			500
// @Success			200
// @Router			/products [get]
func (p *ProductRoutes) GetAllProducts(c *gin.Context) {
	products, err := Models.Product().GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Retrieved all products successfully",
		"data": gin.H{
			"products": products,
		},
	})
}

// @Summary			Get a product in store by ID
// @Description		Get a product in store by ID
// @Tags			Products
// @ID				get-product-by-id
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			200
// @Param			productID	path	uint true "Product ID"
// @Router			/products/{productID} [get]
func (p *ProductRoutes) GetProductByID(c *gin.Context) {
	productIDStr := c.Param("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid product id",
		})
		return
	}

	product, err := Models.Product().GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product retrieved successfully",
		"data": gin.H{
			"product": product,
		},
	})
}

// @Summary			Get a product image
// @Description		Get a product image
// @Tags			Products
// @ID				get-product-image
// @Produce			json
// @Failure			404
// @Success			200
// @Param			imageName	path	string true "Image Filename"
// @Router			/products/images/{imageName} [get]
func (p *ProductRoutes) GetProductImage(c *gin.Context) {
	imageName := c.Param("imageName")

	imagePath := uploadDir + imageName

	// Check if the file exists and is not a directory
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "file does not exist")
		return
	}

	c.File(imagePath)
}

// @Summary			Create a new product
// @Description		Create a new product
// @Tags			Products
// @ID				create-product
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			201
// @Param			name formData string true "Product Name"
// @Param			description formData string true "Product Description"
// @Param			price formData string true "Product Price"
// @Param			image formData file true "Product Image"
// @Param			available formData string true "Product Availability, true or false"
// @Router			/products [post]
func (p *ProductRoutes) CreateProduct(c *gin.Context) {
	var product interfaces.Product

	switch {

	case strings.Contains(c.GetHeader("Content-Type"), "application/x-www-form-urlencoded"):
	case strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data"):
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Failed to parse form",
			})
			return
		}

		priceStr := form.Value["price"]
		name := form.Value["name"]
		description := form.Value["description"]
		image := form.File["image"]
		availableStr := form.Value["available"]

		if priceStr[0] == "" || name[0] == "" || description[0] == "" || image[0] == nil || availableStr[0] == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Product info not complete",
			})
			return
		}

		c.SaveUploadedFile(image[0], fmt.Sprintf("%s/%s", uploadDir, image[0].Filename))

		price, err := strconv.ParseFloat(priceStr[0], 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid price, price must be a number",
			})
			return
		}

		available, err := strconv.ParseBool(availableStr[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Available must be either true or false",
			})
			return
		}

		product = interfaces.Product{
			Name:        name[0],
			Description: description[0],
			Price:       price,
			Image:       imgRoute + image[0].Filename,
			Available:   available,
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "To upload product image, create product via form-data",
			"error": gin.H{
				"error": fmt.Sprintf("You made use of a :%s: header", c.GetHeader("Content-Type")),
			},
		})
		return
	}

	err := Models.Product().CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "An error occured, failed to create product",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Product created successfully",
		"data": gin.H{
			"productID": product.ID,
		},
	})
}

// @Summary			Update an existing product by ID
// @Description		Update an existing product by ID
// @Tags			Products
// @ID				update-product
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			202
// @Param			productID path string true "Product ID"
// @Param			name formData string false "Product Name"
// @Param			description formData string false "Product Description"
// @Param			price formData string false "Product Price"
// @Param			image formData file false "Product Image"
// @Param			available formData string false "Product Availability, true or false"
// @Router			/products/{productID} [put]
func (p *ProductRoutes) UpdateProductByID(c *gin.Context) {
	var product interfaces.Product
	productIDStr := c.Param("productID")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ProductID must be a positive integer",
		})
		return
	}

	switch {

	case strings.Contains(c.GetHeader("Content-Type"), "application/x-www-form-urlencoded"):
	case strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data"):
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Failed to process update form",
			})
			return
		}

		priceStr := form.Value["price"]
		name := form.Value["name"]
		description := form.Value["description"]
		image := form.File["image"]
		availableStr := form.Value["available"]

		if name[0] != "" {
			product.Name = name[0]
		}

		if description[0] != "" {
			product.Description = description[0]
		}

		if image[0] != nil {
			c.SaveUploadedFile(image[0], fmt.Sprintf("%s/%s", uploadDir, image[0].Filename))
			product.Image = imgRoute + image[0].Filename
		}

		var price float64
		if priceStr[0] != "" {
			price, err = strconv.ParseFloat(priceStr[0], 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Invalid price, price must be a number",
				})
				return
			}
			product.Price = price
		}

		var available bool
		if availableStr[0] != "" {
			available, err = strconv.ParseBool(availableStr[0])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Available must be either true or false",
				})
				return
			}
			product.Available = available
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "To upload product image, update product via form-data",
		},
		)
		return
	}

	err = Models.Product().UpdateProductByID(&product, uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "An error occured, couldn't update product",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Product updated successfully",
	})
}

// @Summary			Delete a product in store by ID
// @Description		Delete a product in store by ID
// @Tags			Products
// @ID				delete-product-by-id
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			202
// @Param			product_id	path	uint true "Product ID"
// @Router			/products/{product_id} [delete]
func (p *ProductRoutes) DeleteProductByID(c *gin.Context) {
	productIDStr := c.Param("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid product id",
		})
		return
	}

	err = Models.Product().DeleteProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "An error occured while deleting product",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Product deleted successfully",
	})
}
