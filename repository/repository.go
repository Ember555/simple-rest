package repository

import "simple-rest/models"

type PokemonRepository interface {
	QueryAll() (*[]models.PokemonModel, error)
	SignUp(body *models.UserModel) error
}
