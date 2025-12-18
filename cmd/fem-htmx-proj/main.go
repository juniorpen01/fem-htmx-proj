package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type (
	Count    int
	Contacts []Contact
)

func (c Contacts) hasEmail(email string) bool {
	for _, contact := range c {
		if contact.Email == email {
			return true
		}
	}
	return false
}

const PORT = "localhost:42069"

var mockContacts = Contacts{
	{"idiot", "idiot@gmail.com"},
	{"dummkopf", "dummkopf101@gmail.com"},
}

func main() {
	data := struct {
		Count
		Contacts
	}{0, mockContacts}

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
		name, email := r.FormValue("name"), r.FormValue("email")

		if data.hasEmail(email) {

			http.Error(w, "dupe", http.StatusConflict)
			return
		}

		contact := Contact{name, email}

		data.Contacts = append(data.Contacts, contact)
		tmpl.ExecuteTemplate(w, "contacts", data)
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
