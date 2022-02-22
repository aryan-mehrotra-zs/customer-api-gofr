package customer

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"example.com/customer-api/models"
	"example.com/customer-api/stores"
)

type service struct {
	store stores.Customer
}

func New(store stores.Customer) service {
	return service{store: store}
}

func (s service) Create(ctx *gofr.Context, customer models.Customer) (int64, error) {
	if customer.Name == "" {
		return 0, errors.InvalidParam{Param: []string{"name"}}
	}

	return s.store.Create(ctx, customer)
}

func (s service) Get(ctx *gofr.Context, id string) (models.Customer, error) {
	val, err := strconv.Atoi(id)
	if err != nil || val < 1 {
		return models.Customer{}, errors.InvalidParam{Param: []string{"id"}}
	}

	return s.store.Get(ctx, id)
}

func (s service) Update(ctx *gofr.Context, customer models.Customer) (models.Customer, error) {
	params := make([]string, 0, 2)

	if customer.ID < 1 {
		params = append(params, "id")
	}
	if customer.Name == "" {
		params = append(params, "name")
	}

	if len(params) > 0 {
		return models.Customer{}, errors.InvalidParam{Param: params}
	}

	if err := s.store.Update(ctx, customer); err != nil {
		return models.Customer{}, err
	}

	return s.Get(ctx, strconv.Itoa(customer.ID))

}

func (s service) Delete(ctx *gofr.Context, id string) error {
	val, err := strconv.Atoi(id)
	if err != nil || val < 1 {
		return errors.InvalidParam{Param: []string{"id"}}
	}

	return s.store.Delete(ctx, id)
}
