package repository

import (
	"authPract"
	"github.com/jmoiron/sqlx"
)

type Team interface {
	CreateTeam(userId int, team authPract.Team) (int, error)
	AddUserToTeam(userId int, teamId int) (int, error)
	GetByUserId(userId int) ([]authPract.Team, error)
	GetById(Id int) (authPract.Team, error)
	Delete(projectId int) error
	Update(projectId int, input authPract.Team) (authPract.Team, error)
}

type Repository struct {
	Team
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Team: NewTeamPostgres(db),
	}
}
