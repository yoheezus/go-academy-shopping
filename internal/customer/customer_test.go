package customer

import (
	"testing"
)

// CUST_REQ_001 Once created, a customer should have a unique number that identifies them
func TestCUST_REQ_001(t *testing.T) {
	cust := New()
	if cust.ID == 0 {
		t.Fatal("new customer does not have a unique identifier")
	}
}

// CUST_REQ_002 No two customers should be able to have the same identifer
func TestCUST_REQ_002(t *testing.T) {
	cust := New()
	cust2 := New()
	if cust.ID == cust2.ID {
		t.Fatal("two customers have the same identifier")
	}
}

// CUST_REQ_003 I should not be able to create a new customer that has no email
func TestCUST_REQ_003(t *testing.T) {
	//happy path
	var cust Customer
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
	cust := New()
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
	cust := New()
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
	expected := New()
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

	// check what happends when we try to get a customer with an invalid ID
	returned, err = GetByID(-1)
	if err == nil {
		t.Log("returned:", returned)
		t.Fatal("we got a returned customer, we should have got back an error")
	}

}

// CUST_REQ_010 = After a customer is created, I should be able to update that customer by its unique number
func TestCUST_REQ_010(t *testing.T) {
	original := New()
	original.SetEmail("original@TestCUST_REQ_010.com")
	original.SetFirstName("Bob")
	original.SetLastName("Dole")

	returned, err := GetByID(original.ID)
	if err != nil {
		t.Log("err:", err)
		t.Fatal("Expected customer was not returned, even though it exists")
	}
	// Update returned customers field
	returned.SetFirstName("John")
	returned.SetLastName("Silver")
	returned.SetEmail("altered@testCUST_REQ_010.com")

	updated, err := GetByID(original.ID)
	if err != nil {
		t.Log("err:", err)
		t.Fatal("Expected customer was not returned, even though it exists")
	}
	if updated.FirstName == original.FirstName {
		t.Fatal("Updated first name matches original firstname after updating")
	}

	if updated.LastName == original.LastName {
		t.Fatal("Updated first name matches original firstname after updating")
	}

	if updated.Email != original.Email {
		t.Fatal("It should not be possible to update the email")
	}
}

// CUST_REQ_011 - I should not be able to update a customer's first name to an empty value
func TestCUST_REQ_011(t *testing.T) {
	original := New()
	original.SetEmail("original@TestCUST_REQ_011.com")
	original.SetFirstName("Bob")
	original.SetLastName("Dole")

	returned, err := GetByID(original.ID)
	if err != nil {
		t.Log("err:", err)
		t.Fatal("Expected customer was not returned, even though it exists")
	}
	// Update returned customers field
	err = returned.SetFirstName("")
	if err == nil {
		t.Log("err:", err)
		t.Log("Customer first name was:", original.FirstName)
		t.Log("Customer first name is now:", returned.FirstName)
		t.Fatal("Success on updating firstname to empty when it should be fail.")
	}

}

// CUST_REQ_012 - I should not be able to update a customer's last name to an empty value
func TestCUST_REQ_012(t *testing.T) {
	original := New()
	original.SetEmail("original@TestCUST_REQ_012.com")
	original.SetFirstName("Bob")
	original.SetLastName("Dole")

	returned, err := GetByID(original.ID)
	if err != nil {
		t.Log("err:", err)
		t.Fatal("Expected customer was not returned, even though it exists")
	}
	// Update returned customers field
	err = returned.SetLastName("")
	if err == nil {
		t.Log("err:", err)
		t.Log("Customer last name was:", original.LastName)
		t.Log("Customer last name is now:", returned.LastName)
		t.Fatal("Success on updating lastname to empty when it should be fail.")
	}

}

// CUST_REQ_013 - I should not be able to update a customer's email at all
func TestCUST_REQ_013(t *testing.T) {
	original := New()
	original.SetEmail("original@TestCUST_REQ_013.com")
	original.SetFirstName("Bob")
	original.SetLastName("Dole")

	returned, err := GetByID(original.ID)
	if err != nil {
		t.Log("err:", err)
		t.Fatal("Expected customer was not returned, even though it exists")
	}

	err = returned.SetEmail("altered@TestCUST_REQ_013.com")
	if err == nil {
		t.Fatal("Success on changing email after it was set. This should fail")
	}

	err = returned.SetEmail("")
	if err == nil {
		t.Fatal("Success on changing email to empty after it was set. This should fail")
	}

	if returned.Email != original.Email {
		t.Fatal("Email was changed from the original when this should not be possible.")
	}

}

// CUST_REQ_014 - I should be able to list out all users that have been created
func TestCUST_REQ_014(t *testing.T) {
	// Clear out previous customers
	for k := range customers {
		delete(customers, k)
	}

	cust1 := New()
	cust2 := New()

	all_customers := GetAll()

	if len(all_customers) != 2 {
		t.Log("all customers:", all_customers)
		t.Fatal("customers slice should have 2 values, but it does not.")
	}

	cust3 := New()
	all_customers = GetAll()
	if len(all_customers) != 3 {
		t.Log("all customers:", all_customers)
		t.Fatal("customers slice should have 3 values, but it does not.")
	}

	manual_customers := make([]Customer, 0)
	manual_customers = append(manual_customers, cust1, cust2, cust3)

	for i, c := range all_customers {
		if c.ID != manual_customers[i].ID {
			t.Log("all_customers value", c)
			t.Log("manual_customers value", manual_customers[i])
			t.Fatal("Customer in all_customers does not match customer in manual_customers")
		}
	}
}
