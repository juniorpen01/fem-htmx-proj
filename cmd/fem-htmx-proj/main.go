package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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

	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, count)
	})

	http.HandleFunc("POST /count/{$}", func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		fmt.Fprintf(w, "Count %d", count.Count)
	})

	// start server
	log.Printf("Server started at http://%s", PORT) // the "http://" is for convenience, idk if always correct tho
	http.ListenAndServe(PORT, nil)
}
