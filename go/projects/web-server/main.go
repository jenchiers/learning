package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// the star is a pointer, so pointing to the request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported!", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Golang knows to look for the index.html file in the folder specified below:
	fileServer := http.FileServer(http.Dir("./static"))
	// handle the root route
	http.Handle("/", fileServer)
	// two handlers that link to functions
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// ListenAndServer creates the server, the first param is the port, the second is an error value or nil (when no error is thrown) when created
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// this means there was an error
		log.Fatal(err)
	}
}
