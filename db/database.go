package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=jobs_db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database.\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database.")
	db.Logger = logger.Default.LogMode(logger.Info)

	//TODO: db migrations

	Database = db

}
