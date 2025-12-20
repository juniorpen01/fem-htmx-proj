// Package server
package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type contact struct {
	Name, Email string
}
type contacts []contact

const addr = ":8080"

func Run() {
	// template
	tmpl := template.Must(template.ParseFiles("web/template/index.html"))

	// state
	data := struct {
		Count    int
		Contacts contacts
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
		data.Contacts = append(data.Contacts, newContact)

		tmpl.ExecuteTemplate(w, "contact", newContact)
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
