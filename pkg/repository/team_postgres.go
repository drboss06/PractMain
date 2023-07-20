package repository

import (
	"authPract"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TeamPostgres struct {
	db *sqlx.DB
}

func (r *TeamPostgres) GetByUserId(userId int) ([]authPract.Team, error) {
	var teams []authPract.Team
	query := fmt.Sprintf(`SELECT us.id, us.name, us.team_description as description FROM %s as us join %s as ut on us.id = ut.team_id where ut.user_id = $1`, teamTable, userTeamTable)
	if err := r.db.Select(&teams, query, userId); err != nil {
		return teams, err
	}
	return teams, nil
}

func (r *TeamPostgres) GetById(Id int) (authPract.Team, error) {
	var item authPract.Team
	query := fmt.Sprintf(`SELECT id, name, team_description as description FROM %s WHERE id = $1`, teamTable)
	if err := r.db.Get(&item, query, Id); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TeamPostgres) Delete(projectId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, teamTable)
	_, err := r.db.Exec(query, projectId)
	return err
}

func (r *TeamPostgres) Update(projectId int, input authPract.Team) (authPract.Team, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Description != "" {
		setValues = append(setValues, fmt.Sprintf("team_description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`,
		teamTable, setQuery, argId)
	args = append(args, projectId)
	fmt.Println(query, args)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return authPract.Team{}, err
	}
	return r.GetById(projectId)
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
