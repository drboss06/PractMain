package service

import (
	"authPract"
	"authPract/pkg/repository"
)

type Team interface {
	CreateTeam(userId int, team authPract.Team) (int, error)
	ParseToken(accessToken string) (int, error)
	SendMailToUser(userEmail string) error
	AddUserToTeam(userId int, teamId int) (int, error)
	GetByUserId(userId int) ([]authPract.Team, error)
	GetById(Id int) (authPract.Team, error)
	Delete(projectId int) error
	Update(projectId int, input authPract.Team) error
}

type Service struct {
	Team
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Team: NewTeamService(repos.Team),
	}
}
