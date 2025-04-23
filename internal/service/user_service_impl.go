package service

import (
	"github.com/gmalheiro/golang-playground-txdb/internal/model"
	"github.com/gmalheiro/golang-playground-txdb/internal/repository"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(r repository.UserRepository) UserService {
	return UserServiceImpl{
		repository: r,
	}
}

func (u UserServiceImpl) GetAll() ([]*model.User, error) {
	return u.repository.GetAll()
}
