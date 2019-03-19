package query

import (
	"log"
	"simple-rest/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func initSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/pokemondb")
	// session, err := mgo.Dial("mongodb://simple-mongodb:27017/pokemondb")
	if err != nil {
		panic(err)
	}
	return session
}

func QueryAll() (*[]models.PokemonModel, error) {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	pokemons := []models.PokemonModel{}
	err := c.Find(nil).All(&pokemons)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pokemons, nil
}

func Query(queryString string) (*models.PokemonModel, error) {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	pokemon := models.PokemonModel{}

	err := c.Find(bson.M{"_id": queryString}).One(&pokemon)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pokemon, nil
}

func Insert(body *models.PokemonModel) error {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	err := c.Insert(body)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Update(body *models.PokemonModel) error {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	err := c.UpdateId(body.ID, body)
	if err != nil {
		log.Println("ID :", body.ID, err)
		return err
	}
	return nil
}

func Delete(id string) error {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("pokemons")
	err := c.Remove(bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SignUp(body *models.UserModel) error {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("user")
	err := c.Insert(body)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Signin(username string, password string) error {
	ss := initSession()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB("pokemondb").C("user")
	user := models.UserModel{}

	err := c.Find(bson.M{"username": username, "password": password}).One(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
