package main

import (
	"log"
	"net/http"

	"simple-rest/route"
)

func initRoute() {
	http.HandleFunc("/search", route.SearchByName)
	http.HandleFunc("/create", route.Create)
}

func main() {
	initRoute()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
