package main

import (
	"authPract/pkg/api"
	"authPract/pkg/repository"
	"authPract/pkg/service"
	adder "authPract/pkg/transfer/grpc"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfug(); err != nil {
		logrus.Fatal("error initializing configs: ", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatalf("error occured while running gRPC server: %s", err)
	}

	s := grpc.NewServer()
	server := adder.NewGrpc(services)

	api.RegisterAdderServer(s, server)
	if err := s.Serve(l); err != nil {
		logrus.Fatalf("error occured while running gRPC server: %s", err)
	}
	//handlers := handler.NewHandler(services)
	//srv := new(authPract.Server)
	//
	//if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	//	logrus.Fatalf("error occured while running http server: %s", err)
	//}

}

func initConfug() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
