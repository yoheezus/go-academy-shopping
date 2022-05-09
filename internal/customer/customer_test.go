package customer

import "testing"

// CUST_REQ_001 Once created, a customer should have a unique number that identifies them
func TestCUST_REQ_001(t *testing.T) {
	var cust customer
	cust = New()
	if cust.ID == 0 {
		t.Fatal("new customer does not have a unique identifier")
	}
}

// CUST_REQ_002 No two customers should be able to have the same identifer
func TestCUST_REQ_002(t *testing.T) {
	var cust customer
	cust = New()
	var cust2 customer
	cust2 = New()
	if cust.ID == cust2.ID {
		t.Fatal("two customers have the same identifier")
	}
}

// CUST_REQ_003 I should not be able to create a new customer that has no email
func TestCUST_REQ_003(t *testing.T) {
	//happy path
	var cust customer
	cust = New()
	email := "fred@CUST_REQ_003.com"
	err := cust.SetEmail(email)
	if err != nil {
		t.Fatal("our email provided should be okay, we should not have had an error")
	}

	if cust.GetEmail() != email {
		t.Fatal("email of cust did not match provided")
		t.Log("email:", email)
	}

	//unhappy path
	cust = New()
	err = cust.SetEmail("")
	if err == nil {
		t.Fatal("our email provided was empty, we should have had an error")
	}
}

// CUST_REQ_004 I should not be able to create a new customer that has no first name
func TestCUST_REQ_004(t *testing.T) {
	//happy path
	var cust customer
	cust = New()
	firstName := "fred"
	err := cust.SetFirstName(firstName)
	if err != nil {
		t.Fatal("our firstName provided should be okay, we should not have had an error")
	}

	if cust.GetFirstName() != firstName {
		t.Log("firstName:", firstName)
		t.Log("cust.GetFirstName():", cust.GetFirstName())
		t.Fatal("firstName of cust did not match provided")
	}

	//unhappy path
	cust = New()
	err = cust.SetFirstName("")
	if err == nil {
		t.Fatal("our firstName provided should no be okay, we should have had an error")
	}
}

// CUST_REQ_005 I should not be able to create a new customer that has no last name
func TestCUST_REQ_005(t *testing.T) {
	//happy path
	var cust customer
	cust = New()
	lastName := "bloggs"
	err := cust.SetLastName(lastName)
	if err != nil {
		t.Fatal("our lastName provided should be okay, we should not have had an error")
	}

	if cust.GetLastName() != lastName {
		t.Log("lastName:", lastName)
		t.Fatal("lastName of cust did not match provided")
	}

	//unhappy path
	cust = New()
	err = cust.SetLastName("")
	if err == nil {
		t.Fatal("our lastName provided is empty, we should have had an error")
	}
}

// CUST_REQ_006 I should not be able to create a customer that has the same email as an existing customer
func TestCUST_REQ_006(t *testing.T) {
	cust := New()
	email := "fred.bloggs@CUST_REQ_006.com"
	err := cust.SetEmail(email)
	if err != nil {
		t.Fatal("this email is unique, we should not have been returned an error")
	}

	cust2 := New()
	err = cust2.SetEmail(email)
	if err == nil {
		t.Fatal("this email is not unique, we should have been returned an error")
	}
}

// CUST_REQ_007 I should be able to create a customer that has the same first name as an existing customer
func TestCUST_REQ_007(t *testing.T) {
	cust := New()
	firstName := "fred"
	err := cust.SetFirstName(firstName)
	if err != nil {
		t.Fatal("this first name is unique, we should not have been returned an error")
	}

	cust2 := New()
	err = cust2.SetFirstName(firstName)
	if err != nil {
		t.Fatal("this first name is not unique, we should have been returned an error")
	}
}

// CUST_REQ_008 I should be able to create a customer that has the same last name as an existing customer
func TestCUST_REQ_008(t *testing.T) {
	cust := New()
	lastName := "bloggs"
	err := cust.SetLastName(lastName)
	if err != nil {
		t.Fatal("this last name is unique, we should not have been returned an error")
	}

	cust2 := New()
	err = cust2.SetLastName(lastName)
	if err != nil {
		t.Fatal("this last name is not unique, we should have been returned an error")
	}
}

// CUST_REQ_009 After a customer is created, I should be able to get that customer back by its unique number
func TestCUST_REQ_009(t *testing.T) {
	var expected customer
	expected = New()
	expected.SetEmail("expected@TestCUST_REQ_009.com")
	expected.SetFirstName("terry")
	expected.SetLastName("bloggs")

	returned, err := GetByID(expected.ID)
	if err != nil {
		t.Log("err:", err)
		t.Fatal("the id in expected does not return a customer, we got an error")
	}
	if returned != expected {
		t.Log("returned:", returned)
		t.Log("expected:", expected)
		t.Fatal("returned customer is not the same as the expected")
	}
}
