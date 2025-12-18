package contacts_test

import (
	"slices"
	"testing"

	contact_store "github.com/juniorpen01/fem-htmx-proj/internal"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		input    contact_store.Contact
		expected []contact_store.Contact
	}{
		{contact_store.Contact{"idiot", "idiot@gmail.com"}, []contact_store.Contact{{"idiot", "idiot@gmail.com"}}},
		{contact_store.Contact{"dummkopf", "dummkopf101@gmail.com"}, []contact_store.Contact{{"idiot", "idiot@gmail.com"}, {"dummkopf", "dummkopf101@gmail.com"}}},
		{contact_store.Contact{"therealdummkopf", "dummkopf101@gmail.com"}, []contact_store.Contact{{"idiot", "idiot@gmail.com"}, {"dummkopf", "dummkopf101@gmail.com"}}},
	}

	contacts := contact_store.Contacts{}

	for _, c := range cases {
		contacts.Add(c.input)
		if actual, expected := contacts.Contacts(), c.expected; !slices.Equal(actual, expected) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
