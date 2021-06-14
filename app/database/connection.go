package database

import (
	"fmt"
	"log"

	"github.com/itokun99/bang-jago/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDb() *gorm.DB {
	var db = DB
	var gormConfig gorm.Config
	appConfig := config.GetConfig()

	username := appConfig.Database.DbUsername
	password := appConfig.Database.DbPassword
	database := appConfig.Database.DbName
	host := appConfig.Database.DbHost
	port := appConfig.Database.DbPort
	driver := appConfig.Database.DbDriver

	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local"), &gormConfig)
		if err != nil {
			// fmt.Println("DB err: ", err)
			log.Fatalln(err)
		}
		// Only for debugging
		if err == nil {
			fmt.Println("DB connection successful! ^_^")
		}
		break
	case "postgres":
		db, err = gorm.Open(postgres.Open("host="+host+" port="+port+" user="+username+" dbname="+database+" password="+password), &gormConfig)
		if err != nil {
			// fmt.Println("DB err: ", err)
			log.Fatalln(err)
		}
		// Only for debugging
		if err == nil {
			fmt.Println("DB connection successful! ^_^")
		}
		break
	default:
		log.Fatalln("The driver " + driver + " is not implemented yet T_T")
	}

	DB = db

	return DB
}

// GetDB - get a connection
func GetDB() *gorm.DB {
	return DB
}
