package postgres

import (
	"fmt"
	"github.com/sirupsen/logrus"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"self-payroll-rakamin/config"
)

func DBInit() *gorm.DB {
	config := config.NewConfig()
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.GetHost(), config.GetUser(), config.GetPass(), config.GetDBName(), config.GetPort())
	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatal("Failed Connect To Database, Error : " + err.Error())
	}

	return db

}
