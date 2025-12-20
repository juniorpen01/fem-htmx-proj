// Package server
package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type form struct {
	// Values contact NOTE: doesnt seem to serve a purpose with how i did it
	Error string
}

type contact struct {
	Name, Email string
}
type contacts []contact

func (c *contacts) add(contact contact) error {
	if c.hasEmail(contact.Email) {
		return errors.New("duplicate email")
	} else {
		*c = append(*c, contact)
		return nil
	}
}

func (c contacts) hasEmail(email string) bool {
	for _, contact := range c {
		if email == contact.Email {
			return true
		}
	}
	return false
}

const addr = ":8080"

func Run() {
	// template
	tmpl := template.Must(template.ParseFiles("web/template/index.html"))

	// state
	data := struct {
		Count    int
		Contacts contacts
		Form     form
	}{Contacts: contacts{{"idiot@gmail.com", "idiot@gmail.com"}, {"dummkopf", "dummkopf101@gmail.com"}}}

	// routes

	// serve static files NOTE: kinda copied from chatgpt but ye i think i understand it
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.HandleFunc("POST /count", func(w http.ResponseWriter, r *http.Request) {
		data.Count++
		fmt.Fprint(w, data.Count)
	})

	http.HandleFunc("POST /contacts", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")

		newContact := contact{name, email}
		if err := data.Contacts.add(newContact); err != nil {
			// data.Form.Values = newContact
			data.Form.Error = err.Error()

			w.WriteHeader(http.StatusConflict)
			tmpl.ExecuteTemplate(w, "contact-error", data.Form)
		} else {
			data.Form.Error = ""
			tmpl.ExecuteTemplate(w, "contact", newContact)
		}
	})

	// start
	log.Printf("server started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, logger(http.DefaultServeMux)))
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
