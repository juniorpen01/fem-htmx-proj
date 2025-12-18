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
)

func main() {
	const PORT = "localhost:42069"

	mockContactStore := contact_store.Contacts{}

	mockContactStore.Add(contact_store.Contact{Name: "idiot", Email: "idiot@gmail.com"})
	mockContactStore.Add(contact_store.Contact{Name: "dummkopf", Email: "dummkopf@gmail.com"})
	mockContactStore.Add(contact_store.Contact{Name: "therealdummkopf", Email: "dummkopf@gmail.com"})

	data := struct {
		Count
		Contacts
	}{}

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

		if err := mockContactStore.Add(contact); err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		data.Contacts = mockContactStore.Contacts()

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
