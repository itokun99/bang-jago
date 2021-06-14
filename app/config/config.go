package config

import (
	"fmt"
	"os"

	dotenv "github.com/joho/godotenv"
)

type AppConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func initEnvironment() {
	if err := dotenv.Load(); err != nil {
		fmt.Println("Error initEnvironment()", err)
		panic(err)
	}
}

// function fot get config environtment file
func GetConfig() AppConfig {
	var appConfig AppConfig

	// init dotenv environment
	initEnvironment()

	// init env db
	appConfig.Database.DbName = os.Getenv("DB_NAME")
	appConfig.Database.DbUsername = os.Getenv("DB_USER")
	appConfig.Database.DbPassword = os.Getenv("DB_PASS")
	appConfig.Database.DbHost = os.Getenv("DB_HOST")
	appConfig.Database.DbPort = os.Getenv("DB_PORT")
	appConfig.Database.DbDriver = os.Getenv("DB_DRIVER")

	// init env server
	appConfig.Server.ServerEnv = os.Getenv("APP_ENV")
	appConfig.Server.ServerPort = os.Getenv("APP_PORT")
	appConfig.Server.ServerName = os.Getenv("APP_NAME")
	appConfig.Server.ServerKey = os.Getenv("APP_KEY")

	return appConfig
}
