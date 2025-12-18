package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	contact_store "github.com/juniorpen01/fem-htmx-proj/internal"
)

type (
	Count    int
	Contacts []contact_store.Contact
	Form     struct {
		Values contact_store.Contact
		Error  string
	}
)

func main() {
	const PORT = "localhost:42069"

	contactStore := contact_store.Contacts{}

	contactStore.Add(contact_store.Contact{Name: "idiot", Email: "idiot@gmail.com"})
	contactStore.Add(contact_store.Contact{Name: "dummkopf", Email: "dummkopf@gmail.com"})
	contactStore.Add(contact_store.Contact{Name: "therealdummkopf", Email: "dummkopf@gmail.com"})

	data := struct {
		Count
		Contacts
		Form
	}{Contacts: contactStore.Contacts()}

	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	router.HandleFunc("POST /count/{$}", func(w http.ResponseWriter, r *http.Request) {
		data.Count++
		tmpl.ExecuteTemplate(w, "count", data)
	})

	router.HandleFunc("POST /contacts/{$}", func(w http.ResponseWriter, r *http.Request) {
		contact := contact_store.Contact{Name: r.FormValue("name"), Email: r.FormValue("email")}

		if err := contactStore.Add(contact); err != nil {
			data.Values.Name = contact.Name
			data.Values.Email = contact.Email
			data.Error = err.Error()
			w.WriteHeader(http.StatusConflict)
			tmpl.ExecuteTemplate(w, "contact-form", data)
			return
		} else {
			data.Contacts = contactStore.Contacts()

			tmpl.ExecuteTemplate(w, "contact-form", data)
			tmpl.ExecuteTemplate(w, "oob-contact", contact)
		}
	})

	// start server
	log.Printf("Server started at http://%s", PORT) // the "http://" is for convenience, idk if always correct tho
	http.ListenAndServe(PORT, logging(router))
}

// copied from the dreamsofcode guy, manually typed tho; ok i understand this now
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Println(r.Method, r.URL.Path, time.Since(start))
		})
}
