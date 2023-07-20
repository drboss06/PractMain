package service

import (
	"authPract"
	"authPract/pkg/repository"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"time"
)

const (
	salt       = "asdjkalhsd123123laksj"
	signingKey = "kaijdhOAS;KD'JJAKsjd"
	tokenTTL   = 12 * time.Hour
)

type TeamService struct {
	repo repository.Team
}

func (s *TeamService) GetByUserId(userId int) ([]authPract.Team, error) {
	return s.repo.GetByUserId(userId)
}

func (s *TeamService) GetById(Id int) (authPract.Team, error) {
	return s.repo.GetById(Id)
}

func (s *TeamService) Delete(projectId int) error {
	return s.repo.Delete(projectId)
}

func (s *TeamService) Update(projectId int, input authPract.Team) error {
	return s.repo.Update(projectId, input)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) CreateTeam(userId int, team authPract.Team) (int, error) {
	return s.repo.CreateTeam(userId, team)
}

func (s *TeamService) SendMailToUser(userEmail string) error {
	from := "s4sha228@yandex.ru"

	user := "s4sha228@yandex.ru"
	password := "sfbeguevylkbnyva"

	to := []string{
		userEmail,
	}

	addr := "smtp.yandex.ru:587"
	host := "smtp.yandex.ru"

	msg := []byte("Subject: Test mail\r\n\r\n" +
		"localhost:8000/team/add\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	//fmt.Println("Email sent successfully")
	return nil
}

func (s *TeamService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("couldn't parse claims")
	}
	return claims.UserID, nil
}

func (s *TeamService) AddUserToTeam(userId int, teamId int) (int, error) {
	return s.repo.AddUserToTeam(userId, teamId)
}
