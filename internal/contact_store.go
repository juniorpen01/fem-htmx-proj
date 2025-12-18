// Package contacts is some next level overengineering dont do this unless ur experienced
package contacts

import "errors"

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Contacts struct {
	contacts []Contact
}

func (c Contacts) Contacts() []Contact {
	return c.contacts
}

func (c *Contacts) Add(contact Contact) error {
	switch {
	case contact.Name == "":
		return errors.New("no name")
	case contact.Email == "":
		return errors.New("no email")
	case c.hasEmail(contact.Email):
		return errors.New("duplicate email")
	default:
		c.contacts = append(c.contacts, contact)
		return nil
	}
}

func (c Contacts) hasEmail(email string) bool {
	for _, contact := range c.contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}
