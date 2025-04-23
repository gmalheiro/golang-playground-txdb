package repository

import "github.com/gmalheiro/golang-playground-txdb/internal/model"

type UserRepository interface {
	GetAll() ([]*model.User, error)
}
