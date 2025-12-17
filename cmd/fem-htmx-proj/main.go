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

const PORT = "localhost:42069"

func main() {
	count := Count{}

	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", count)
	})

	router.HandleFunc("POST /count/{$}", func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		tmpl.ExecuteTemplate(w, "count", count)
	})

	// start server
	log.Printf("Server started at http://%s", PORT) // the "http://" is for convenience, idk if always correct tho
	http.ListenAndServe(PORT, logging(router))
}

// copied from the dreamsofcode guy, manually typed tho
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Println(r.Method, r.URL.Path, time.Since(start))
		})
}
