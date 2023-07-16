package repository

import (
	"authPract"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user authPract.User) (int, error)
	GetUser(username, password string) (authPract.User, error)
	DeleteUser(id int) error
	GetUserById(id int) (authPract.UserWidthPassHash, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
