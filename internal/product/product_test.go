package product

import "testing"

// PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them
func TestPRODUCT_REQ_001(t *testing.T) {
	var p product
	p = New()

	// 0 is the default value for an int
	if p.ID == 0 {
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}
}

// PRODUCT_REQ_002 No two products should be able to have the same identifer
func TestPRODUCT_REQ_002(t *testing.T) {
	p1 := New()
	p2 := New()
	if p1.ID == p2.ID {
		t.Fatal("PRODUCT_REQ_002: No two products should be able to have the same identifer")
	}
}

// PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product
func TestPRODUCT_REQ_003(t *testing.T) {
	p1 := New()
	productName := "whatever"
	err := p1.SetName(productName)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_003: I should not be able to create a new product that has the same name as an existing product")
	}

	p2 := New()
	err = p2.SetName(productName)
	if err == nil {
		t.Fatal("PRODUCT_REQ_003: I should not be able to create a new product that has the same name as an existing product")
	}
}

// PRODUCT_REQ_004 I should be able to create a new product that has the same price as an existing product
func TestPRODUCT_REQ_004(t *testing.T) {
	p1 := New()
	productPrice := 100
	err := p1.SetPrice(productPrice)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_004: I should be able to create a new product that has the same price as an existing product")
	}

	p2 := New()
	err = p2.SetPrice(productPrice)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_004: I should be able to create a new product that has the same price as an existing product")
	}
}

// PRODUCT_REQ_005 I should be able to create a new product that has the same stock amount as an existing product
func TestPRODUCT_REQ_005(t *testing.T) {
	p1 := New()
	stockAmount := 5
	err := p1.SetStock(stockAmount)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_005: I should be able to create a new product that has the same stock amount as an existing product")
	}

	p2 := New()
	err = p2.SetStock(stockAmount)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_005: I should be able to create a new product that has the same stock amount as an existing product")
	}
}
