package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"self-payroll-rakamin/config"
	"self-payroll-rakamin/database/postgres"
	"self-payroll-rakamin/migration"
)

var (
	cfgFile, asciiFile *string
)

func init() {
	cfgFile = flag.String("c", "./config/app.json", "Configuration")
	asciiFile = flag.String("A", "haradakumiko.txt", "Show Banner")
}

func printASCIIArtFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {

	if err := config.InitializeConfig(*cfgFile); err != nil {
		logrus.Fatal("Error reading configuration file : " + err.Error())
	}

	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	postgresDatabase := postgres.DBInit()

	migration.Migrate(postgresDatabase)

	server.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<img  style=\"display: block; margin: auto;\" git  src=\"https://preview.redd.it/e6i2fgwn1ng71.jpg?width=640&crop=smart&auto=webp&s=8eea2b644d2b00bd5110343bdeea1f75be3daf4c\" />")
	})

	err := printASCIIArtFromFile(*asciiFile)

	fmt.Printf("apps running on %s\n", color.Green(config.GetString("server.port")))
	err = server.Start(config.GetString("server.port"))

	if err != nil {
		logrus.Fatal("Error : " + err.Error())
	}
}
