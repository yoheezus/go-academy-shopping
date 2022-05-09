package customer

import "errors"

var nextIdentifier int

func New() customer {
	nextIdentifier++
	return customer{
		ID: nextIdentifier,
	}
}

func (c *customer) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be blank")
	}
	c.Email = email
	return nil
}

func (c *customer) GetEmail() string {
	return c.Email
}

func (c *customer) SetFirstName(firstName string) error {
	if firstName == "" {
		return errors.New("first name cannot be blank")
	}
	c.FirstName = firstName
	return nil
}

func (c *customer) GetFirstName() string {
	return c.FirstName
}
