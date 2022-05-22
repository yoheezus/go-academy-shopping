package order

import (
	"customer"
	"product"
	"testing"
)

var (
	test_prod     = product.New("Test Product 1")
	test_customer = customer.New()
)

// ORDER_REQ_001 - A customer that has sufficient funds should be able to order a product
func TestORDER_REQ_001(t *testing.T) {
	test_order, err := New(test_customer, test_prod)
	if err != nil {

	}
}

// ORDER_REQ_002 - A customer that does not have sufficient funds should not be able to order a product
func TestORDER_REQ_002(t *testing.T) {

}

// ORDER_REQ_003 - A customer that has ordered a product should have the price of that product deducted from their funds
func TestORDER_REQ_003(t *testing.T) {

}
