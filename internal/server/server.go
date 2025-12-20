// Package server
package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const addr = ":8080"

func Run() {
	// template
	tmpl := template.Must(template.ParseFiles("web/template/index.html"))

	// state
	count := struct{ Count int }{}

	// routes
	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, count)
	})

	http.HandleFunc("POST /count", func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		fmt.Fprintf(w, "count %d", count.Count)
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
