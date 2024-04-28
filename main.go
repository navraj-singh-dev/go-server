package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Remember that this handler is called by the form.html itself.
		From the browser you must open form.html and submit it
		then this handler will run, otherise from browser it cannot
		be run successfully, error will come if not called from
		form.html
	*/

	// return error response if request came on wrong route
	if r.URL.Path != "/form" {
		http.Error(w, "Wrong route", http.StatusBadGateway)
		return
	}

	// parse the form, well form came in the request must be parsed before getting data from it
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Could not parse the form that was in the request:\n%v", err)
		return
	}

	// Fill the response with success message and data
	fmt.Fprintf(w, "POST request successfull\n\n")
	// extract data from form
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	message := r.Form.Get("message")
	// add extracted data to response
	fmt.Fprintf(w, "Name = %v\n", name)
	fmt.Fprintf(w, "Email = %v\n", email)
	fmt.Fprintf(w, "Message = %v\n", message)

}

func greethandler(w http.ResponseWriter, r *http.Request) {
	// return error if wrong endpoint
	if r.URL.Path != "/greet" {
		http.Error(w, "Wrong path, send request to /greetings", http.StatusBadRequest)
		return
	}

	// return error if wrong HTTP method
	if r.Method != "GET" {
		http.Error(w, "Wrong method, Use GET", http.StatusMethodNotAllowed)
		return
	}

	// good, now serve the data in response
	fmt.Fprintf(w, "Hello, You Made A Successful Request\nCongratulations!")

}

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
