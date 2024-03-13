package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(writer http.ResponseWriter, reader *http.Request) {
	if err := reader.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm Err %v", err)
		return
	}
	fmt.Fprintf(writer, "POST request accepted")
	name := reader.FormValue("name")
	address := reader.FormValue("address")
	fmt.Fprintf(writer, "Received Name is %v", name)
	fmt.Fprintf(writer, "Received Address is %v", address)
}

func helloHandler(writer http.ResponseWriter, reader *http.Request) {
	// validating if correct request is routed
	if reader.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "Helllo hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
