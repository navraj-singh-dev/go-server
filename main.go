package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// define the port to listen
	PORT := ":8000"

	// ------------ static files setup --------
	// just because of this line of code, static files can be served
	// very important line
	fileServer := http.FileServer(http.Dir("./static/"))

	// ------------ routes ----------------
	// this route serves the index.html and form.html static files
	http.Handle("/", fileServer)

	// this route takes the data from "form" and sends it as response
	http.HandleFunc("/form", formHandler)

	// this route sends a greet message
	http.HandleFunc("/greet", greethandler)

	// ---------- listen on the server -----------
	fmt.Printf("Server is UPPP.. and running on PORT  %s\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatalln(err)
	}
}
