package models

// PokemonModel is
type PokemonModel struct {
	ID      string `bson:"_id" json:"_id"`
	Name    string `bson:"name" json:"name"`
	Element string `bson:"element" json:"element"`
	Weight  string `bson:"weight" json:"weight"`
}
