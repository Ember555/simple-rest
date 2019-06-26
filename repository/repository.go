package repository

import "simple-rest/models"

type PokemonRepository interface {
	QueryAll() (*[]models.PokemonModel, error)
	Query(queryString string) (*models.PokemonModel, error)
	Insert(body *models.PokemonModel) error
	Update(body *models.PokemonModel) error
	Delete(id string) error
}

type UserRepository interface {
	SignUp(body *models.UserModel) error
	Signin(username string, password string) error
}
