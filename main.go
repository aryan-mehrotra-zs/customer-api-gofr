package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	handlers "example.com/customer-api/handlers/customer"
	services "example.com/customer-api/services/customer"
	stores "example.com/customer-api/stores/customer"
)

func main() {
	app := gofr.New()

	store := stores.New()
	service := services.New(store)
	h := handlers.New(service)

	app.POST("/customer", h.Create)
	app.GET("/customer/{id}", h.Get)
	app.PUT("/customer", h.Update)
	app.DELETE("/customer/{id}", h.Delete)

	app.Start()

}
