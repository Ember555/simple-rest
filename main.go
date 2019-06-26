package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-rest/handler/pokemon"
	"simple-rest/models"
	"simple-rest/repository"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	r := mux.NewRouter()

	config := models.DBConfig{
		Host:          "127.0.0.1:27017", //simple-mongodb:27017 for docker
		Username:      "mongodb",
		Pass:          "password",
		DB:            "pokemondb",
		PokeColletion: "pokemons",
		UserColletion: "user",
	}
	uri := fmt.Sprintf("%v://%v/%v", config.Username, config.Host, config.DB)
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Println(err)
		return
	}
	pokemonRepository := repository.NewPokemonRepo(session, config.DB, config.PokeColletion)
	userRepository := repository.NewUserRepo(session, config.DB, config.UserColletion)

	pokemonHandler, err2 := pokemon.NewHandler(pokemonRepository, userRepository)
	if err2 != nil {
		log.Println(err)
		return
	}

	pokemonHandler.Handle(r)
	// r.HandleFunc("/search", route.SearchAll).Methods("GET")
	// r.HandleFunc("/search/id/{id}", route.SearchByName).Methods("GET")
	// r.HandleFunc("/create", route.Create).Methods("POST")
	// r.HandleFunc("/update", route.Update).Methods("POST")
	// r.HandleFunc("/delete/id/{id}", route.Delete).Methods("POST")

	// r.HandleFunc("/signup", route.Signup).Methods("POST")
	// r.HandleFunc("/signin", route.Signin).Methods("POST")

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8085", r))
}
