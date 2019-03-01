package query

import (
	"log"
	"simple-rest/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func initSession() *mgo.Session {
	// session, err := mgo.Dial("mongodb://127.0.0.1:27017/pokemondb")
	session, err := mgo.Dial("mongodb://simple-mongodb:27017/pokemondb")
	if err != nil {
		panic(err)
	}
	return session
}

func QueryAll() *[]models.PokemonModel {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	pokemons := []models.PokemonModel{}
	err := c.Find(nil).All(&pokemons)
	if err != nil {
		log.Fatal(err)
	}
	return &pokemons
}

func Query(name string) *models.PokemonModel {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	pokemon := models.PokemonModel{}
	err := c.Find(bson.M{"name": name}).One(&pokemon)
	if err != nil {
		log.Fatal(err)
	}
	return &pokemon
}
