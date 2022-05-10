package customer

import "errors"

var (
	nextIdentifier int
	AllEmails      map[string]bool
)

func init() {
	AllEmails = make(map[string]bool)
}

func New() customer {
	nextIdentifier++
	newCustomer := customer{ID: nextIdentifier}
	save(&newCustomer)
	return newCustomer
}

func (c *customer) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be blank")
	}

	if checkIfEmailInUse(email) {
		return errors.New("email cannot be a duplicate")
	}

	c.Email = email
	updateInUseEmail(email)
	save(c)
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
	save(c)
	return nil
}

func (c *customer) GetFirstName() string {
	return c.FirstName
}

func (c *customer) SetLastName(lastName string) error {
	if lastName == "" {
		return errors.New("last name cannot be blank")
	}
	c.LastName = lastName
	save(c)
	return nil
}

func (c *customer) GetLastName() string {
	return c.LastName
}

func checkIfEmailInUse(email string) bool {
	return AllEmails[email]
}

func updateInUseEmail(email string) {
	AllEmails[email] = true
}

func GetByID(id int) (customer, error) {
	return restore(id)
}
