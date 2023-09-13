package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Job struct {
	Id          int       `gorm:"id"`
	Company     string    `gorm:"company"`
	JobTitle    string    `gorm:"job_title"`
	DateApplied time.Time `gorm:"date_applied"`
	JobUrl      string    `gorm:"job_url"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		dsn := fmt.Sprintf(
			"host=jobs_db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect database.")
		}

		var job Job

		db.First(&job, 1)

		str := job.JobTitle

		return c.SendString(str)
	})

	app.Listen(":3000")

}
