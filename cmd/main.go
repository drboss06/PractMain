package main

import (
	"authPract/pkg/transfer/http"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfug(); err != nil {
		logrus.Fatal("error initializing configs: ", err)
	}

	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	Password: viper.GetString("db.password"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//})
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db: %s", err)
	//}
	//
	//repos := repository.NewRepository(db)
	//services := service.NewService(repos)
	//handlers := handler.NewHandler(services)
	//srv := new(authPract.Server)
	//
	//l, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	logrus.Fatalf("error occured while running gRPC server: %s", err)
	//}
	//
	//s := grpc.NewServer()
	//server := adder.NewGrpc(services)
	//server := &adder.GRPCServer{}
	//api.RegisterAdderServer(s, server)

	//transHttp.RunRest()
	http.RunRest()
	//httpProxy := http.NewHttpProxy(services)
	//httpProxy.RunRest()

	//if err := s.Serve(l); err != nil {
	//	logrus.Fatalf("error occured while running gRPC server: %s", err)
	//}
	//
	//if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	//	logrus.Fatalf("error occured while running http server: %s", err)
	//}

}

//func unaryInterceptor(
//	ctx context.Context,
//	req interface{},
//	info *grpc.UnaryServerInfo,
//	handler grpc.UnaryHandler,
//) (interface{}, error) {
//	log.Println("--> unary interceptor: ", info.FullMethod)
//	return handler(ctx, req)
//}

func initConfug() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
