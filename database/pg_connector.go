package database

import (
	"fmt"
	"log"

	"gin-api/models"
	"gin-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var user, password, dbName, dbHost, dsn, port string

const (
	keyPGUser     = "PGUSER"
	keyPGPassword = "PGPASSWORD"
	keyPGDatabase = "PGDATABASE"
	keyPGHost     = "PGHOST"
	keyPGPort     = "PGPORT"
)

func init() {
	user = utils.GetValueFromEnv(keyPGUser)
	password = utils.GetValueFromEnv(keyPGPassword)
	dbName = utils.GetValueFromEnv(keyPGDatabase)
	dbHost = utils.GetValueFromEnv(keyPGHost)
	port = utils.GetValueFromEnv(keyPGPort)
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, user, password, dbName, port)

	var err error
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to the Database: ", err)
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	DB = database
}
