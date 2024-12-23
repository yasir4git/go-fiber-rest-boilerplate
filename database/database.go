package database

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/models"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/config"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.AppConfig.DB_PORT
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.AppConfig.DB_HOST, port, config.AppConfig.DB_USER, config.AppConfig.DB_PASSWORD, config.AppConfig.DB_NAME)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database Connected successfully")

	// TODO: always drop table just for development
	err = DB.Migrator().DropTable(&models.User{})
	if err != nil {
		panic("failed to drop table")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error AutoMigrate database: %v", err)
	}
}
