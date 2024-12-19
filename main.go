package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Static file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Route handlers
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/new", newHandler)
	mux.HandleFunc("/create", createHandler)
	mux.HandleFunc("/edit", editHandler)
	mux.HandleFunc("/update", updateHandler)
	mux.HandleFunc("/delete", deleteHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
