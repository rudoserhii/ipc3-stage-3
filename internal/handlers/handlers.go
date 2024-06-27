package handlers

import "github.com/obiMadu/ipc3-stage-3/internal/interfaces"

type Handlers struct {
	product ProductRoutes
	order   OrderRoutes
}

var Models interfaces.Models

func NewHandlers(models interfaces.Models) *Handlers {
	Models = models
	return &Handlers{
		product: ProductRoutes{},
		order:   OrderRoutes{},
	}
}

func (h *Handlers) Product() interfaces.ProductRoutes {
	return &h.product
}

func (h *Handlers) Order() interfaces.OrderRoutes {
	return &h.order
}
