package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
)

type OrderRoutes struct{}

// @Summary			Get all orders in store
// @Description		Get all orders in store
// @Tags			Orders
// @ID				get-all-orders
// @Produce			json
// @Failure			500
// @Success			200
// @Router			/orders [get]
func (p *OrderRoutes) GetAllOrders(c *gin.Context) {
	orders, err := Models.Order().GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve orders",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Retrieved all orders successfully",
		"data": gin.H{
			"orders": orders,
		},
	})
}

// @Summary			Retrieve an order by ID
// @Description		Retrieve an order by ID
// @Tags			Orders
// @ID				get-order-by-id
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			200
// @Param			orderID	path	uint true "Order ID"
// @Router			/products/{orderID} [get]
func (p *OrderRoutes) GetOrderByID(c *gin.Context) {
	orderIDStr := c.Param("orderID")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid order id",
		})
		return
	}

	order, err := Models.Order().GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Order retrieved successfully",
		"data": gin.H{
			"order": order,
		},
	})
}

// @Summary			Create a new order
// @Description		Create a new order
// @Tags			Orders
// @ID				create-order
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			201
// @Param			productID formData uint true "Product ID"
// @Param			quantity formData uint true "Product Quantity to place order for"
// @Router			/orders [post]
func (p *OrderRoutes) CreateOrder(c *gin.Context) {
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

		productIDStr := form.Value["productID"][0]
		quantityStr := form.Value["quantity"][0]

		if productIDStr == "" || quantityStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "You must specify a product and a quantity to place an order",
			})
			return
		}

		productID, err := strconv.ParseUint(productIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "ProductID must be a positive integer",
			})
			return
		}

		quantity, err := strconv.ParseUint(quantityStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Quantity must be a positive integer",
			})
			return
		}

		order := &interfaces.Order{
			ProductID: uint(productID),
			Quantity:  uint(quantity),
		}

		err = Models.Order().CreateOrder(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "An error occured, failed to create order",
				"error": gin.H{
					"error": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Order created successfully",
			"data": gin.H{
				"orderID": order.ID,
			},
		})

	}
}

// @Summary			Cancel an order
// @Description		Cancel an order
// @Tags			Orders
// @ID				cancel-order
// @Produce			json
// @Failure			400
// @Failure			500
// @Success			201
// @Param			orderID formData string true "Order ID"
// @Router			/orders/{orderID} [delete]
func (p *OrderRoutes) CancelOrderByID(c *gin.Context) {
	orderIDStr := c.Param("orderID")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid order id",
		})
		return
	}

	err = Models.Order().CancelOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to cancel order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Order cancelled successfully",
	})
}
