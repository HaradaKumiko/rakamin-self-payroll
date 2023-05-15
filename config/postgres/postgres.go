package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"self-payroll-rakamin/model/domain"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "")
	if err != nil {
		logrus.Fatal("Failed Connect To Database, Error : " + err.Error())
	}
	db.AutoMigrate(&domain.Position{})
	return db
}
