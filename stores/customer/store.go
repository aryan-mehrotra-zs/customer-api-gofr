package customer

import (
	"database/sql"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"example.com/customer-api/models"
	"example.com/customer-api/stores"
)

type store struct{}

func New() stores.Customer {
	return store{}
}

func (s store) Create(ctx *gofr.Context, customer models.Customer) (int64, error) {
	var id int64

	err := ctx.DB().QueryRowContext(ctx, create, customer.Name).Scan(&id)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	return id, nil
}

func (s store) Get(ctx *gofr.Context, id string) (models.Customer, error) {
	var customer models.Customer

	db := ctx.DB()
	err := db.QueryRowContext(ctx, get, id).Scan(&customer.ID, &customer.Name)

	switch {
	case err == sql.ErrNoRows:
		return models.Customer{}, errors.EntityNotFound{Entity: "customer", ID: id}
	case err != nil:
		return models.Customer{}, errors.DB{Err: err}
	}

	return customer, nil
}

func (s store) Update(ctx *gofr.Context, customer models.Customer) error {
	_, err := ctx.DB().ExecContext(ctx, update, customer.Name, customer.ID)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (s store) Delete(ctx *gofr.Context, id string) error {
	_, err := ctx.DB().ExecContext(ctx, delete, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
