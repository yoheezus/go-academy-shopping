package order

import (
	"github.com/go-academy-shopping/internal/customer"
	"github.com/go-academy-shopping/internal/product"
)

func New(customer *customer.Customer, product *product.Product) Order {
	return Order{
		Customer: customer,
		Product:  product,
	}

}

func (o *Order) Process() error {
	return nil
}
