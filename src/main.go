package main

// go run .\src\main.go

import (
	"fmt"
	"log"
	"net/http"
)

var handlerWelcome = func(w http.ResponseWriter, r *http.Request) {
	shouldReturn := validate(r, w, "/welcome", "GET")
	if !shouldReturn {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the web-server..."))
}

var handleForm = func(w http.ResponseWriter, r *http.Request) {
	shouldReturn := validate(r, w, "/form", "POST")
	if !shouldReturn {
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseFrom err : %v", err)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name : %s", name)
	fmt.Fprintf(w, "Email : %s", email)
}

func validate(r *http.Request, w http.ResponseWriter, endpoint string, method string) bool {
	if r.URL.Path != endpoint {
		http.Error(w, "404 not found", http.StatusNotFound)
		return false
	}

	if r.Method != method {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func main() {
	log.Println("Starting web server.")

	fileServer := http.FileServer(http.Dir("static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/welcome", handlerWelcome)

	if errors := http.ListenAndServe(":8080", nil); errors != nil {
		log.Fatalf("%v", errors)
	}
}
