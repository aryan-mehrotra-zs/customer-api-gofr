package customer

import (
	"strings"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"example.com/customer-api/models"
	"example.com/customer-api/services"
)

type handler struct {
	service services.Customer
}

func New(service services.Customer) handler {
	return handler{service: service}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var customer models.Customer

	if err := ctx.Bind(&customer); err != nil {
		return int64(0), errors.InvalidParam{Param: []string{"body"}}
	}

	id, err := h.service.Create(ctx, customer)
	if err != nil {
		return int64(0), err
	}

	return id, nil
}

func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
	id := strings.TrimSpace(ctx.PathParam("id"))

	customer, err := h.service.Get(ctx, id)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var customer models.Customer

	if err := ctx.Bind(&customer); err != nil {
		return models.Customer{}, errors.InvalidParam{Param: []string{"body"}}
	}

	customer, err := h.service.Update(ctx, customer)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := strings.TrimSpace(ctx.PathParam("id"))

	err := h.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
