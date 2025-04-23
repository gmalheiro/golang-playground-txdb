package repository

import (
	"database/sql"
	"github.com/gmalheiro/golang-playground-txdb/internal/model"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) GetAll() ([]*model.User, error) {
	rows, err := u.db.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var u *model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
