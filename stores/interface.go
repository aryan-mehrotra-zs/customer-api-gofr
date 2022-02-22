package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"example.com/customer-api/models"
)

type Customer interface {
	Create(ctx *gofr.Context, customer models.Customer) (int64, error)
	Get(ctx *gofr.Context, id string) (models.Customer, error)
	Update(ctx *gofr.Context, customer models.Customer) error
	Delete(ctx *gofr.Context, id string) error
}
