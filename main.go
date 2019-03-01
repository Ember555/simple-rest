package main

import (
	"log"
	"net/http"

	"simple-rest/route"
)

func initRoute() {
	http.HandleFunc("/search", route.SearchAll)
	http.HandleFunc("/search/", route.SearchByName)
}

func main() {
	initRoute()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
