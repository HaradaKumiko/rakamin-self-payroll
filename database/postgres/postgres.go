package postgres

import (
	"fmt"
	"github.com/sirupsen/logrus"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"self-payroll-rakamin/config"
)

func DBInit() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.GetString("credential.database_postgres_host"),
		config.GetString("credential.database_postgres_user"),
		config.GetString("credential.database_postgres_password"),
		config.GetString("credential.database_postgres_database"),
		config.GetString("credential.database_postgres_port"),
	)
	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatal("Failed Connect To Database, Error : " + err.Error())
	}

	return db

}
