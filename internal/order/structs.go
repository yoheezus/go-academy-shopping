package order

import (
	"github.com/go-academy-shopping/internal/customer"
	"github.com/go-academy-shopping/internal/product"
)

type Order struct {
	ID       int
	Customer *customer.Customer
	Product  *product.Product
	Complete bool
}
