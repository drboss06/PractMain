package service

import (
	"authPract"
	"authPract/pkg/repository"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
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

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) ParseCSV(fileName string) (string, error) {
	c := make(chan int)
	defer timer("main")()

	var db = repository.DbClick{}

	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	var csvMet = CsvProp{}
	var reader = csvMet.parceCsvFile(viper.GetString("filedir") + fileName)
	var tableName = fileName
	var dbName = viper.GetString("clickdb.database")

	fmt.Println("Create queries...")

	createQuery, insertQueries := getSQL(reader, dbName, tableName)
	_, err = conn.Query(ctx, createQuery)

	fmt.Println("Create table...")

	if err != nil && err.Error() != "EOF" && err.Error() != "code: 20, message: Number of columns doesn't match" {
		fmt.Println(createQuery)
		log.Fatal(err.Error())
	}

	fmt.Println("Insert data...")

	//for _, query := range insertQueries {
	//	if _, err := conn.Query(ctx, query); err != nil && err.Error() != "EOF" {
	//		log.Fatal(err.Error())
	//	}
	//}

	//for _, query := range insertQueries {
	//	if err := conn.AsyncInsert(ctx, query, false); err != nil && err.Error() != "EOF" {
	//		log.Fatal(err.Error())
	//	}
	//}
	//go insert_to_db(insertQueries, ctx, conn, c)

	for i := 0; i < len(insertQueries); i++ {
		go insert_to_db(insertQueries[i], ctx, conn, c, i)
	}

	for i := 0; i < len(insertQueries); i++ {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "func is ready")
	}

	fmt.Println("All right")

	return fileName, nil
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

func initConfug() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
