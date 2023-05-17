package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func InitializeConfig(filename string) error {
	splits := strings.Split(filepath.Base(filename), ".")

	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(filename))
	viper.SetConfigType(splits[1])

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error : " + err.Error())
		return err
	}
	return nil
}

func isSet(key string) {
	if !viper.IsSet(key) {
		logrus.Fatal("Configuration key %s not found", key)
		os.Exit(1)
	}
}
func GetString(key string) string {
	isSet(key)
	return viper.GetString(key)
}

func GetInt(key string) int {
	isSet(key)
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	isSet(key)
	return viper.GetBool(key)
}
