package service

import (
	"authPract"
	"authPract/pkg/repository"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
