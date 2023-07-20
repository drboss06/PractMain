package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/spf13/viper"
)

type DbClick struct {
}

func (d *DbClick) Connect() (driver.Conn, error) {

	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("error init configs: ")
	//}
	if err := initConfig(); err != nil {
		return nil, err
	}

	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{viper.GetString("clickdb.host") + ":" + viper.GetString("clickdb.port")},
			Auth: clickhouse.Auth{
				Database: viper.GetString("clickdb.database"),
				Username: viper.GetString("clickdb.usrname"),
				Password: viper.GetString("clickdb.password"),
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			//TLS: &tls.Config{
			//	InsecureSkipVerify: true,
			//},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}

	return conn, nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
