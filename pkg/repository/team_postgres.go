package repository

import (
	"authPract"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TeamPostgres struct {
	db *sqlx.DB
}

func NewTeamPostgres(db *sqlx.DB) *TeamPostgres {
	return &TeamPostgres{db: db}
}

func (r *TeamPostgres) CreateTeam(userId int, team authPract.Team) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTeamQuery := fmt.Sprintf("INSERT INTO %s (name, team_description) VALUES ($1, $2) RETURNING id", teamTable)
	row := tx.QueryRow(createTeamQuery, team.Name, team.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersTeamQuery := fmt.Sprintf("INSERT INTO %s (user_id, team_id) VALUES ($1, $2)", userTeamTable)
	_, err = tx.Exec(createUsersTeamQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TeamPostgres) AddUserToTeam(userId int, teamId int) (int, error) {
	var id int
	addUserToTeamQuery := fmt.Sprintf("INSERT INTO %s (team_id, user_id) VALUES ($1, $2) RETURNING id", teamUserTable)

	row := r.db.QueryRow(addUserToTeamQuery, teamId, userId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
