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
		t.Fatal("firstName of cust did not match provided")
		t.Log("firstName:", firstName)
	}

	//unhappy path
	cust = New()
	err = cust.SetFirstName("")
	if err == nil {
		t.Fatal("our firstName provided should no be okay, we should have had an error")
	}
}
