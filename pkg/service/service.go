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
}

type Service struct {
	Team
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Team: NewTeamService(repos.Team),
	}
}
