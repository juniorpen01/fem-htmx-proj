package main

import (
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

	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("web/template/index.html")
		if err != nil {
			log.Fatalln(err)
		}

		tmpl.Execute(w, count)

		count.Count++
	})

	// start server
	log.Printf("Server started at http://%s", PORT) // the "http://" is for convenience, idk if always correct tho
	http.ListenAndServe(PORT, nil)
}
