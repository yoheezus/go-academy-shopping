package product

import "testing"

// PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them
func TestPRODUCT_REQ_001(t *testing.T) {
	var p product
	p = New()

	// 0 is the default value for an int
	if p.GetID() == 0 {
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}

	p2 := New()
	if p2.GetID() != (p.GetID() + 1) {
		t.Log("p.GetID():", p.GetID())
		t.Log("p2.GetID():", p2.GetID())
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}

	p3 := New()
	if p3.GetID() != (p2.GetID() + 1) {
		t.Log("p2.GetID():", p2.GetID())
		t.Log("p3.GetID():", p3.GetID())
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}
}

// PRODUCT_REQ_002 No two products should be able to have the same identifer
func TestPRODUCT_REQ_002(t *testing.T) {
	p1 := New()
	p2 := New()
	if p1.GetID() == p2.GetID() {
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
	returnedName := p1.GetName()
	if productName != returnedName {
		t.Log("productName:", productName)
		t.Log("returnedName:", returnedName)
		t.Fatal("PRODUCT_REQ_003: I should not be able to create a new product that has the same name as an existing product")
	}

	p2 := New()
	err = p2.SetName(productName)
	if err == nil {
		t.Fatal("PRODUCT_REQ_003: I should not be able to create a new product that has the same name as an existing product")
	}
	returnedName = p2.GetName()
	expectedName := ""
	if returnedName == productName {
		t.Log("returnedName:", returnedName)
		t.Fatal("PRODUCT_REQ_003: I should not be able to create a new product that has the same name as an existing product")
	}
	if returnedName != expectedName {
		t.Log("returnedName:", returnedName)
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
	if productPrice != p1.GetPrice() {
		t.Log("productPrice:", productPrice)
		t.Log(" p1.GetPrice():", p1.GetPrice())
		t.Fatal("PRODUCT_REQ_004: I should be able to create a new product that has the same price as an existing product")
	}

	p2 := New()
	err = p2.SetPrice(productPrice)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_004: I should be able to create a new product that has the same price as an existing product")
	}
	if productPrice != p2.GetPrice() {
		t.Log("productPrice:", productPrice)
		t.Log(" p2.GetPrice():", p2.GetPrice())
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
	if stockAmount != p1.GetStock() {
		t.Log("stockAmount:", stockAmount)
		t.Log(" p1.GetStock():", p1.GetStock())
		t.Fatal("PRODUCT_REQ_005: I should be able to create a new product that has the same stock amount as an existing product")
	}

	p2 := New()
	err = p2.SetStock(stockAmount)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_005: I should be able to create a new product that has the same stock amount as an existing product")
	}
	if stockAmount != p2.GetStock() {
		t.Log("stockAmount:", stockAmount)
		t.Log(" p1.GetStock():", p2.GetStock())
		t.Fatal("PRODUCT_REQ_005: I should be able to create a new product that has the same stock amount as an existing product")
	}
}

// PRODUCT_REQ_006 Once created, I should be able to get a product back by its unique identifier
func TestPRODUCT_REQ_006(t *testing.T) {
	expected := New()
	expectedName := "coffee"
	expectedPrice := 350
	expectedStock := 200
	expected.SetName(expectedName)
	expected.SetPrice(expectedPrice)
	expected.SetStock(expectedStock)

	returned, err := GetByID(expected.GetID())
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_006: Once created, I should be able to get a product back by its unique identifier")
	}

	if returned.GetName() != expectedName {
		t.Log("returned.GetName():", returned.GetName())
		t.Log("expectedName:", expectedName)
		t.Fatal("PRODUCT_REQ_006: Once created, I should be able to get a product back by its unique identifier")
	}
	if returned.GetPrice() != expectedPrice {
		t.Log("returned.GetPrice():", returned.GetPrice())
		t.Log("expectedPrice:", expectedPrice)
		t.Fatal("PRODUCT_REQ_006: Once created, I should be able to get a product back by its unique identifier")
	}
	if returned.GetStock() != expectedStock {
		t.Log("returned.GetStock():", returned.GetStock())
		t.Log("expectedStock:", expectedStock)
		t.Fatal("PRODUCT_REQ_006: Once created, I should be able to get a product back by its unique identifier")
	}
}
