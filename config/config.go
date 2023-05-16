package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type (
	Config struct {
	}
)

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error... :" + err.Error())
	}
	return &Config{}
}

func (c Config) GetHost() string {
	return os.Getenv("DB_HOST")
}

func (c Config) GetDBName() string {
	return os.Getenv("DB_NAME")
}

func (c Config) GetPort() string {
	return os.Getenv("DB_PORT")
}

func (c Config) GetUser() string {
	return os.Getenv("DB_USER")
}

func (c Config) GetPass() string {
	return os.Getenv("DB_PASS")
}
