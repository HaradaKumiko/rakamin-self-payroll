package migration

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"self-payroll-rakamin/app/domain/models"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Position{},
	)

	if err != nil {
		logrus.Fatal("Something wrong, Error : ", err)
	}
}
