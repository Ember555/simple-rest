package models

// PokemonModel is
type PokemonModel struct {
	ID      string
	Name    string `bson:"name"`
	Element string
	Weight  string
}
