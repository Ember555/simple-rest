package main

import (
	"log"
	"net/http"

	"simple-rest/route"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", route.SearchAll).Methods("GET")
	r.HandleFunc("/search/id/{id}", route.SearchByName).Methods("GET")
	r.HandleFunc("/create", route.Create).Methods("POST")
	r.HandleFunc("/update", route.Update).Methods("POST")
	r.HandleFunc("/delete/id/{id}", route.Delete).Methods("POST")

	r.HandleFunc("/signup", route.Signup).Methods("POST")
	r.HandleFunc("/signin", route.Signin).Methods("POST")

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8085", r))
}
