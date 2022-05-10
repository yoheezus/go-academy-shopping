package product

import "testing"

// PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them
func TestPRODUCT_REQ_001(t *testing.T) {
	var p product
	p = New()
	if p.GetId() == 0 {
		t.Fatal("PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them")
	}

	p2 := New()
	if p2.GetId() != (p.GetId() + 1) {
		t.Log("p.GetId():", p.GetId())
		t.Log("p2.GetId():", p2.GetId())
		t.Fatal("PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them")
	}

	p3 := New()
	if p3.GetId() != (p2.GetId() + 1) {
		t.Log("p2.GetId():", p2.GetId())
		t.Log("p3.GetId():", p3.GetId())
		t.Fatal("PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them")
	}
}

// PRODUCT_REQ_002 No two products should be able to have the same identifer
func TestPRODUCT_REQ_002(t *testing.T) {
	p1 := New()
	p2 := New()

	if p1.ID == p2.ID {
		t.Fatal("PRODUCT_REQ_002 No two products should be able to have the same identifer")
	}

}

// PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product
func TestPRODUCT_REQ_003(t *testing.T) {
	p1 := New()
	productName := "whatever"
	err := p1.SetName(productName)

	if err != nil {
		t.Fatal("PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product")
	}

	returnedName := p1.GetName()
	if productName != returnedName {
		t.Log("productName:", productName)
		t.Log("returnedName:", returnedName)
		t.Fatal("PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product")
	}

	p2 := New()
	err = p2.SetName(productName)
	expectedName := ""
	if err == nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product")
	}
	returnedName = p2.GetName()
	if returnedName == productName {
		t.Log("returnedName:", returnedName)
		t.Fatal("PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product")
	}
	if returnedName != expectedName {
		t.Fatal("PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product")
	}
}

// PRODUCT_REQ_004 I should be able to create a new product that has the same price as an existing product
func TestPRODUCT_REQ_004(t *testing.T) {
	p1 := New()
	productPrice := 100
	err := p1.SetPrice(productPrice)

	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_004 I should be able to create a new product that has the same price as an existing product")
	}

	p2 := New()
	err = p2.SetPrice(productPrice)

	if err != nil {
		t.Fatal("PRODUCT_REQ_004 I should be able to create a new product that has the same price as an existing product")
	}
}

// PRODUCT_REQ_005 I should be able to create a new product that has the same stock amount as an existing product
func TestPRODUCT_REQ_005(t *testing.T) {
	p1 := New()
	stockAmt := 100
	err := p1.SetStock(stockAmt)

	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_005 I should be able to create a new product that has the same stock amount as an existing product")
	}

	p2 := New()
	err = p2.SetStock(stockAmt)

	if err != nil {
		t.Fatal("PRODUCT_REQ_005 I should be able to create a new product that has the same stock amount as an existing product")
	}
}

// PRODUCT_REQ_006 Once created, I should be able to get a product back by its unique identifier
func TestPRODUCT_REQ_006(t *testing.T) {
	expected := New()
	expected.SetName("coffee")
	expected.SetPrice(350)
	expected.SetStock(200)

	returned, err := GetByID(expected.ID)

	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_006 Once created, I should be able to get a product back by its unique identifier")
	}

	if expected != returned {
		t.Log("expected:", expected)
		t.Log("returned:", returned)
		t.Fatal("PRODUCT_REQ_006 Once created, I should be able to get a product back by its unique identifier")
	}
}
