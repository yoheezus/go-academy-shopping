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

// PRODUCT_REQ_007 Once created, I should be able to update a product's price
func TestPRODUCT_REQ_007(t *testing.T) {
	p1 := New()
	initialPrice := 200
	err := p1.SetPrice(initialPrice)
	if err != nil {
		t.Fatal("price could not be properly set")
	}

	if p1.GetPrice() != initialPrice {
		t.Log("p1.GetPrice():", p1.GetPrice())
		t.Fatal("price has not been set correctly")
	}

	updatedPrice := 500
	err = p1.SetPrice(updatedPrice)

	if p1.GetPrice() != updatedPrice {
		t.Log("p1.GetPrice():", p1.GetPrice())
		t.Log("updatedPrice:", updatedPrice)
		t.Fatal("price was not updated")
	}
}

// PRODUCT_REQ_008 Once created, I should be able to update a product's stock amount
func TestPRODUCT_REQ_008(t *testing.T) {
	p1 := New()
	initialStock := 200
	err := p1.SetStock(initialStock)
	if err != nil {
		t.Fatal("stock could not be properly set")
	}

	if p1.GetStock() != initialStock {
		t.Log("p1.GetStock():", p1.GetStock())
		t.Fatal("stock has not been set correctly")
	}

	updatedStock := 500
	err = p1.SetStock(updatedStock)

	if p1.GetStock() != updatedStock {
		t.Log("p1.GetStock():", p1.GetStock())
		t.Log("updatedStock:", updatedStock)
		t.Fatal("stock was not updated")
	}

	updatedStock = -1
	err = p1.SetStock(updatedStock)
	if err == nil {
		t.Fatal("cannot set stock to below 0")
	}
}

// PRODUCT_REQ_009 Once created, I should be able to put a product on sale, so it's returned price should be a percentage of its full price
func TestPRODUCT_REQ_009(t *testing.T) {
	p1 := New()
	initialPrice := 300
	p1.SetPrice(initialPrice)
	t.Log("Initial Percentage:", p1.SalePercentage)
	discount := 50
	err := p1.SetSalePercentage(discount)
	discountedPrice := 150

	if err != nil {
		t.Fatal("was not able to discount product")
	}

	if p1.GetPrice() != discountedPrice {
		t.Log("discount:", discount)
		t.Log("p1.SalePercentage", p1.SalePercentage)
		t.Log("p1.GetPrice():", p1.GetPrice())
		t.Log("discountedPrice", discountedPrice)
		t.Fatal("price is not equal to discountedPrice after discount")
	}

	_ = p1.SetSalePercentage(0)
	if p1.GetPrice() != initialPrice {
		t.Fatal("price did not reset to original")
	}

}

// PRODUCT_REQ_010 I should not be able to create a new product that has a negative stock
func TestPRODUCT_REQ_010(t *testing.T) {
	p1 := New()
	err := p1.SetStock(-50)

	if err == nil {
		t.Fatal("stock cannot be set to a negative number")
	}
}

// PRODUCT_REQ_011 I should not be able to create a new product that has a negative price
func TestPRODUCT_REQ_011(t *testing.T) {
	p1 := New()
	err := p1.SetPrice(-50)

	if err == nil {
		t.Fatal("price cannot be set to a negative number")
	}
}

// PRODUCT_REQ_012 I should not be able to create a new product that has a sale percentage under 90% off
func TestPRODUCT_REQ_012(t *testing.T) {
	p1 := New()
	err := p1.SetSalePercentage(99)

	if err == nil {
		t.Fatal("cannot set sale percentage above 90%")
	}

}
