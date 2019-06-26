package repository

import (
	"log"
	"simple-rest/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userRepo struct {
	conn       *mgo.Session
	db         string
	collection string
}

func NewUserRepo(conn *mgo.Session, db string, collection string) UserRepository {
	return &userRepo{
		conn:       conn,
		db:         db,
		collection: collection,
	}
}

func (p *userRepo) SignUp(body *models.UserModel) error {
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

func (p *userRepo) Signin(username string, password string) error {
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
