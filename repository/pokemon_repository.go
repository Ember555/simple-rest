package repository

import (
	"log"
	"simple-rest/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type pokemonRepo struct {
	conn       *mgo.Session
	db         string
	collection string
}

func NewPokemonRepo(conn *mgo.Session, db string, collection string) PokemonRepository {
	return &pokemonRepo{
		conn:       conn,
		db:         db,
		collection: collection,
	}
}

func (p *pokemonRepo) QueryAll() (*[]models.PokemonModel, error) {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	pokemons := []models.PokemonModel{}
	err := c.Find(nil).All(&pokemons)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pokemons, nil
}

func (p *pokemonRepo) Query(queryString string) (*models.PokemonModel, error) {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	pokemon := models.PokemonModel{}

	err := c.Find(bson.M{"_id": queryString}).One(&pokemon)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pokemon, nil
}

func (p *pokemonRepo) Insert(body *models.PokemonModel) error {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	err := c.Insert(body)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *pokemonRepo) Update(body *models.PokemonModel) error {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	err := c.UpdateId(body.ID, body)
	if err != nil {
		log.Println("ID :", body.ID, err)
		return err
	}
	return nil
}

func (p *pokemonRepo) Delete(id string) error {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	err := c.Remove(bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *pokemonRepo) SignUp(body *models.UserModel) error {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	err := c.Insert(body)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *pokemonRepo) Signin(username string, password string) error {
	ss := p.conn.Copy()
	defer ss.Close()
	ss.SetMode(mgo.Monotonic, true)

	c := ss.DB(p.db).C(p.collection)
	user := models.UserModel{}

	err := c.Find(bson.M{"username": username, "password": password}).One(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
