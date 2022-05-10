package customer

import "errors"

// Map of customers id to a customer pointer
var customers map[int]*customer

func init() {
	customers = make(map[int]*customer)
}

func save(c *customer) error {
	if c.ID == 0 { // dot notation automatically de-references pointers
		return errors.New("id cannot be 0")
	}
	customers[c.ID] = c
	return nil
}

func restore(id int) (customer, error) {
	c, exists := customers[id]
	if !exists {
		return customer{}, errors.New("customer with that id doesn't exist")
	}
	return *c, nil
}
