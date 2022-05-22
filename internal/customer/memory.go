package customer

import "errors"

// map of customers id to a customer pointer
var customers map[int]*Customer

func init() {
	customers = make(map[int]*Customer)
}

func save(c *Customer) error {
	if c.ID == 0 {
		return errors.New("id cannot be 0")
	}
	customers[c.ID] = c
	return nil
}

func restore(id int) (Customer, error) {
	c, exists := customers[id]
	if !exists {
		return Customer{}, errors.New("customer with that id does not exist")
	}
	return *c, nil
}
