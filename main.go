package main

import (
	"net/http"
	"fmt"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Access -> localhost:3000/api")
}

func api(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "request method GET"}`))
		case "POST":
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"message": "request method POST"}`))
		case "PUT":
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(`{"message": "request method PUT"}`))
		case "DELETE":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "request method DELETE"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/api", api)
	log.Fatal(http.ListenAndServe(":3000", nil))
}