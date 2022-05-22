package product

import "testing"

var (
	productName    string
	productPrice   int
	productStock   int
	salePercentage int

	test_prod1 = product{
		ID:             999,
		ProductName:    productName,
		ProductPrice:   productPrice,
		StockAmount:    productStock,
		SalePercentage: salePercentage,
	}
)

// PRODUCT_REQ_001 Once created, a product should have a unique number that identifies them
func TestPRODUCT_REQ_001(t *testing.T) {
	var p product
	p = New("PRODUCT_REQ_001")

	// 0 is the default value for an int
	if p.GetID() == 0 {
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}

	p2 := New("PRODUCT_REQ_001_2")
	if p2.GetID() != (p.GetID() + 1) {
		t.Log("p.GetID():", p.GetID())
		t.Log("p2.GetID():", p2.GetID())
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}

	p3 := New("PRODUCT_REQ_001_3")
	if p3.GetID() != (p2.GetID() + 1) {
		t.Log("p2.GetID():", p2.GetID())
		t.Log("p3.GetID():", p3.GetID())
		t.Fatal("PRODUCT_REQ_001: Once created, a product should have a unique number that identifies them")
	}
}

// PRODUCT_REQ_002 No two products should be able to have the same identifer
func TestPRODUCT_REQ_002(t *testing.T) {
	p1 := New("PRODUCT_REQ_002")
	p2 := New("PRODUCT_REQ_002_3")
	if p1.GetID() == p2.GetID() {
		t.Fatal("PRODUCT_REQ_002: No two products should be able to have the same identifer")
	}
}

// PRODUCT_REQ_003 I should not be able to create a new product that has the same name as an existing product
func TestPRODUCT_REQ_003(t *testing.T) {
	productName := "PRODUCT_REQ_003"
	p1 := New(productName)
	err := p1.SetName(productName)
	if err != nil {
		t.Log(err)
		t.Fatal("PRODUCT_REQ_003: Failed to set the name of a product to it's existing value")
	}

	p2 := New(productName)
	if p2.ProductName != "Faulty Product" {
		t.Fatal("PRODUCT_REQ_003: I should not be able to create a new product that has the same name as an existing product")
	}
	returnedName := p2.GetName()
	expectedName := "Faulty Product"
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
	p1 := New("PRODUCT_REQ_004")
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

	p2 := New("PRODUCT_REQ_004_2")
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
	p1 := New("PRODUCT_REQ_005")
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

	p2 := New("PRODUCT_REQ_005_2")
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
	expected := New("TestPRODUCT_REQ_006")
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
	p1 := New("PRODUCT_REQ_007")
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
	p1 := New("PRODUCT_REQ_008")
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
	productPrice = 300
	test_prod1.SetPrice(productPrice)
	t.Log("Initial Percentage:", test_prod1.SalePercentage)
	salePercentage = 50
	err := test_prod1.SetSalePercentage(salePercentage)
	discountedPrice := 150

	if err != nil {
		t.Fatal("was not able to discount product")
	}

	if test_prod1.GetPrice() != discountedPrice {
		t.Log("discount:", salePercentage)
		t.Log("test_prod1.SalePercentage", test_prod1.SalePercentage)
		t.Log("test_prod1.GetPrice():", test_prod1.GetPrice())
		t.Log("discountedPrice", discountedPrice)
		t.Fatal("price is not equal to discountedPrice after discount")
	}

	_ = test_prod1.SetSalePercentage(0)
	if test_prod1.GetPrice() != productPrice {
		t.Fatal("price did not reset to original")
	}

}

// PRODUCT_REQ_010 I should not be able to create a new product that has a negative stock
func TestPRODUCT_REQ_010(t *testing.T) {
	productStock = -50
	err := test_prod1.SetStock(productStock)

	if err == nil {
		t.Fatal("stock cannot be set to a negative number")
	}
}

// PRODUCT_REQ_011 I should not be able to create a new product that has a negative price
func TestPRODUCT_REQ_011(t *testing.T) {
	productPrice = -50
	err := test_prod1.SetPrice(productPrice)

	if err == nil {
		t.Fatal("price cannot be set to a negative number")
	}
}

// PRODUCT_REQ_012 I should not be able to create a new product that has a sale percentage under 90% off
func TestPRODUCT_REQ_012(t *testing.T) {
	salePercentage = 99
	err := test_prod1.SetSalePercentage(salePercentage)

	if err == nil {
		t.Fatal("cannot set sale percentage above 90%")
	}

}

// PRODUCT_REQ_013 - I should not be able to create a new product that has an empty name
func TestPRODUCT_REQ_013(t *testing.T) {
	productName = "Test Product"
	test_prod1.SetName(productName)

	err := test_prod1.SetName("")

	if err == nil {
		t.Fatal("Should not be able to set product name to blank")
	}

}

// PRODUCT_REQ_014 - I should not be able to update a product to a negative stock
func TestPRODUCT_REQ_014(t *testing.T) {
	productStock = -10
	err := test_prod1.SetStock(productStock)

	if err == nil {
		t.Fatal("Updating a product to negative stock did not throw an errpr")
	}

	if test_prod1.GetStock() == productStock {
		t.Fatal("Stock amount of test_prod1 was set to a negative amount")
	}
}

// PRODUCT_REQ_015 - I should not be able to update a product to a negative price
func TestPRODUCT_REQ_015(t *testing.T) {
	productPrice = -10
	err := test_prod1.SetPrice(productPrice)
	if err == nil {
		t.Fatal("Setting price to negative value did not return error")
	}

	if test_prod1.GetPrice() == productPrice {
		t.Fatal("Price of test_prod1 was set to a negative amount")
	}

}

// PRODUCT_REQ_016 - I should not be able to update a product to a sale percentage under 90% off
func TestPRODUCT_REQ_016(t *testing.T) {
	salePercentage = 91
	productPrice = 100
	err := test_prod1.SetSalePercentage(salePercentage)
	if err == nil {
		t.Fatal("Should not be able to update salesPercentage over 90")
	}
	test_prod1.SetPrice(productPrice)
	if test_prod1.GetPrice() == 10 {
		t.Fatal("SalesPercentage was set to a value of 90. It should not be posible.")
	}

}

// PRODUCT_REQ_017 - I should not be able to update a product to a empty name
func TestPRODUCT_REQ_017(t *testing.T) {
	productName = ""
	new_prod := New("PRODUCT_REQ_017")
	err := new_prod.SetName(productName)
	if err == nil {
		t.Fatal("Should error when name is set to empty")
	}

	if new_prod.GetName() == "" {
		t.Fatal("Product name updated to empty, this should not be possible.")
	}
}

// PRODUCT_REQ_018 - I should be able to list all products
func TestPRODUCT_REQ_018(t *testing.T) {
	for k := range allProducts {
		delete(allProducts, k)
	}

	_ = New("Prod 1")
	_ = New("Prod 2")
	if len(allProducts) != 2 {
		t.Log("count of products:", len(allProducts))
		t.Log("conent of allProducts:", allProducts)
		t.Fatal("allProducts should be set to 2 because there are 2 products.")
	}

	p3 := New("Prod 3")
	if _, exists := allProducts[p3.ID]; !exists || len(allProducts) != 3 {
		t.Fatal("Prod 3 not found in allProducts list or the length of allProducts is not 3")
	}

}

// PRODUCT_REQ_019 - Updates to product fields should persist
func TestPRODUCT_REQ_019(t *testing.T) {
	productPrice = 100
	test_prod1.SetPrice(productPrice)

	returned_prod, err := GetByID(999)
	if err != nil {
		t.Log(err)
		t.Fatal("Not able to return product 999 which exists")
	}

	if returned_prod.GetPrice() != test_prod1.ProductPrice {
		t.Fatal("Price of product was not updated after it had been set in a copy.")
	}

}
