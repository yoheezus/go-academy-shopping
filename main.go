package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-academy-shopping/internal/customer"
)

func main() {

	// customer 1
	cust := customer.New()
	cust.SetEmail("test@test.com")
	cust.SetFirstName("Bobs")
	cust.SetLastName("Uncle")

	//customer 2
	cust2 := customer.New()
	cust2.SetEmail("test@test1.com")
	cust2.SetFirstName("Fred")
	cust2.SetLastName("Borne")

	router := gin.Default()
	router.GET("/test", teapot)
	router.GET("/customers/all", GetAllCustomers)
	router.GET("/customers", GetCustomerById)
	router.POST("/customers", createCustomer)

	router.Run("0.0.0.0:8080")
}

func GetAllCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customer.GetAll())
}

func createCustomer(c *gin.Context) {
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	email := c.PostForm("email")

	fmt.Printf("first_name: %s; last_name: %s; email: %s", firstName, lastName, email)
	// Write Queries to customer
	cust := customer.New()
	cust.SetFirstName(firstName)
	cust.SetLastName(lastName)
	cust.SetEmail(email)
}

func GetCustomerById(c *gin.Context) {
	var id int

	id, _ = strconv.Atoi(c.Query("id"))
	cust, err := customer.GetByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, cust)
	} else {
		c.IndentedJSON(http.StatusCreated, cust)
	}

}

func teapot(c *gin.Context) {
	c.IndentedJSON(http.StatusTeapot, "Teapot")
}
