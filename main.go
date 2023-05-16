package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"self-payroll-rakamin/database/postgres"
	"self-payroll-rakamin/migration"
)

//type (
//	server struct {
//		httpServer *echo.Echo
//		cfg        config.Config
//	}
//
//	Server interface {
//		Run()
//	}
//)
//
//func InitServer(cfg config.Config) Server {
//	e := echo.New()
//
//	return &server{
//		httpServer: e,
//		cfg:        cfg,
//	}
//}
//
//func (s *server) Run() {
//	viper.SetConfigType("json")
//	viper.AddConfigPath("./config")
//	viper.SetConfigName("app")
//
//	err := viper.ReadInConfig()
//	if err != nil {
//		logrus.Fatal("Error : " + err.Error())
//	}
//
//	err = s.httpServer.Start(viper.GetString("server.port"))
//	if err != nil {
//		logrus.Fatal("Error : " + err.Error())
//	}
//}
//
//func main() {
//	config := config.NewConfig()
//	server := InitServer(config)
//
//	wg := sync.WaitGroup{}
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		server.Run()
//	}()
//
//	wg.Wait()
//
//}

func main() {

	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error : " + err.Error())
	}

	server := echo.New()

	postgresDatabase := postgres.DBInit()

	migration.Migrate(postgresDatabase)

	err = server.Start(viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal("Error : " + err.Error())
	}
}
