package customer

import "errors"

var (
	nextIdentifier int
	usedEmails     map[string]bool
)

func init() {
	usedEmails = make(map[string]bool)
}

func New() customer {
	nextIdentifier++
	c := customer{
		ID: nextIdentifier,
	}
	save(&c)
	return c
}

func GetByID(id int) (customer, error) {
	return restore(id)
}

func (c *customer) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be blank")
	}
	if checkIfEmailInUse(email) {
		return errors.New("email already used")
	}
	addInUseEmail(email)
	c.Email = email
	save(c)
	return nil
}

func checkIfEmailInUse(email string) bool {
	return usedEmails[email]
}

func addInUseEmail(email string) {
	usedEmails[email] = true
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
