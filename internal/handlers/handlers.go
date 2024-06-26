package handlers

import "github.com/obiMadu/ipc3-stage-3/internal/interfaces"

type Handlers struct {
	Models  interfaces.Models
	product ProductRoutes
	order   OrderRoutes
}

func NewHandlers(models interfaces.Models) *Handlers {
	return &Handlers{
		Models:  models,
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
