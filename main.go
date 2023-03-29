package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/mrcheese":
		fmt.Fprint(w, "Simon")
	default:
		fmt.Fprint(w, "ERROR!!!")
	}

	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did *not* timeout")
}

func main() {
	http.HandleFunc("/", htmlVsPlain)
	http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var myMux http.ServeMux
	server.Handler = &myMux
	myMux.HandleFunc("/", htmlVsPlain)

	server.ListenAndServe()
}
