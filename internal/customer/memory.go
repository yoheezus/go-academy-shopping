package customer

import "errors"

// map of customers id to a customer pointer
var customers map[int]*customer

func init() {
	customers = make(map[int]*customer)
}

func save(c *customer) error {
	if c.ID == 0 {
		return errors.New("id cannot be 0")
	}
	customers[c.ID] = c
	return nil
}

func restore(id int) (customer, error) {
	c, exists := customers[id]
	if !exists {
		return *c, errors.New("customer with that id does not exist")
	}
	return *c, nil
}
