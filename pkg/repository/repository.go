package repository

import (
	"authPract"
	"github.com/jmoiron/sqlx"
)

type Team interface {
	CreateTeam(userId int, team authPract.Team) (int, error)
	AddUserToTeam(userId int, teamId int) error
}

type Repository struct {
	Team
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Team: NewTeamPostgres(db),
	}
}
