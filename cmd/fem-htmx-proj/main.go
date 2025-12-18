package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Count struct {
	Count int
}

type Contact struct {
	Name, Email string
}

const PORT = "localhost:42069"

func main() {
	count := Count{}

	mockContacts := []Contact{
		{"idiot", "idiot@gmail.com"},
		{"dummkopf", "dummkopf@gmail.com"},
	}

	data := struct {
		Count
		Contacts []Contact
	}{count, mockContacts}

	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	router.HandleFunc("POST /count/{$}", func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		tmpl.ExecuteTemplate(w, "count", count)
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
