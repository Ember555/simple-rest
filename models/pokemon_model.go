package models

// PokemonModel is
type PokemonModel struct {
	ID      string `bson:"_id" json:"_id"`
	Name    string `bson:"name" json:"name"`
	Element string `bson:"element" json:"element"`
	Weight  string `bson:"weight" json:"weight"`
}

type UserModel struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type DBConfig struct {
	Host          string
	Username      string
	Pass          string
	DB            string
	PokeColletion string
	UserColletion string
}
