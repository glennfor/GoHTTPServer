package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PORT = 8000
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
	}
	fmt.Fprint(w, "<h1>HELLO WOLRD!!!</h1>")
}
func form(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "<h1>POST request sucessful</h1><br>")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "<h3>Name: %s</h3><br>", name)
	fmt.Fprintf(w, "<h3>Address: %s</h3><br>", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", form)
	http.HandleFunc("/hello", hello)

	port := fmt.Sprintf(":%v", PORT)

	fmt.Println("Starting server on port", PORT, "....")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
