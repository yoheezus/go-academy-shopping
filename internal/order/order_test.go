package order

import (
	"testing"

	"github.com/go-academy-shopping/internal/customer"
	"github.com/go-academy-shopping/internal/product"
)

var (
	test_prod     = product.New("Test Product 1")
	test_customer = customer.New()
)

// ORDER_REQ_001 - A customer that has sufficient funds should be able to order a product
func TestORDER_REQ_001(t *testing.T) {
	test_customer.Funds = 200
	test_prod.SetPrice(100)
	test_order := New(&test_customer, &test_prod)
	if test_order.Complete {
		t.Fatal("Order just created and as not been processed")
	}

	err := test_order.Process()
	if err != nil {
		t.Log(test_order)
		t.Fatal("Customer had enough funds, Order should have been created")
	}

	if test_customer.Funds != 100 {
		t.Log("test_customer funds:", test_customer.Funds)
		t.Fatal("Funds were not subtracted after order was made")
	}
}

// ORDER_REQ_002 - A customer that does not have sufficient funds should not be able to order a product
func TestORDER_REQ_002(t *testing.T) {
	test_customer.Funds = 50
	test_prod.SetPrice(100)
	test_order := New(&test_customer, &test_prod)
	err := test_order.Process()
	if err == nil {
		t.Log(test_order)
		t.Fatal("Customer did not have enough funds, order should not be possible")
	}

	if test_customer.Funds != 50 {
		t.Log("test_customer funds:", test_customer.Funds)
		t.Fatal("Order went through even though customer did not have enough money")
	}
}

// ORDER_REQ_003 - A customer that has ordered a product should have the price of that product deducted from their funds
func TestORDER_REQ_003(t *testing.T) {

}
