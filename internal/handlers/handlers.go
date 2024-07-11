package handlers

import "github.com/obiMadu/ipc3-stage-3/internal/interfaces"

type Handlers struct {
	product interfaces.ProductRoutes
	order   interfaces.OrderRoutes
}

var Models interfaces.Models

func (h *Handlers) Product() interfaces.ProductRoutes {
	return h.product
}

func (h *Handlers) Order() interfaces.OrderRoutes {
	return h.order
}

func NewHandlers(models interfaces.Models) interfaces.Handlers {
	Models = models
	return &Handlers{
		product: &ProductRoutes{},
		order:   &OrderRoutes{},
	}
}
