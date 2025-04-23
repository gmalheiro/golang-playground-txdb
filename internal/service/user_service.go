package service

import "github.com/gmalheiro/golang-playground-txdb/internal/model"

type UserService interface {
	GetAll() ([]*model.User, error)
}
