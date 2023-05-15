package config

import (
	"github.com/jinzhu/gorm"
	"self-payroll-rakamin/config/postgres"
)

type (
	config struct {
	}

	Config interface {
		Database() *gorm.DB
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.DBInit()
}
