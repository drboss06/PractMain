package service

import (
	"authPract"
	"authPract/pkg/repository"
)

type TeamService struct {
	repo repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) CreateTeam(userId int, team authPract.Team) (int, error) {
	return s.repo.CreateTeam(userId, team)
}
