package routes

import (
	"github.com/esalavat/gojobsapi/db"
	"github.com/esalavat/gojobsapi/db/models"
	"github.com/gofiber/fiber/v2"
)

// type Job struct {
// 	Id          int       `json:"id"`
// 	Company     string    `json:"company"`
// 	JobTitle    string    `json:"job_title"`
// 	DateApplied time.Time `json:"date_applied"`
// 	JobUrl      string    `json:"job_url"`
// }

// func CreateResponseJob(jobModel models.Job) Job {
// 	return Job{
// 		Id:          jobModel.Id,
// 		Company:     jobModel.Company,
// 		JobTitle:    jobModel.JobTitle,
// 		DateApplied: jobModel.DateApplied,
// 		JobUrl:      jobModel.JobUrl,
// 	}
// }

func CreateJob(c *fiber.Ctx) error {
	var job models.Job

	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Database.Create(&job)

	return c.Status(200).JSON(&job)
}

func GetJobs(c *fiber.Ctx) error {
	jobs := []models.Job{}

	db.Database.Find(&jobs)

	return c.Status(200).JSON(jobs)
}

func GetJob(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Could not parse :id param.")
	}

	var job models.Job

	db.Database.Find(&job, "id = ?", id)

	if job.Id == 0 {
		return c.Status(400).JSON("User does not exist.")
	}

	return c.Status(400).JSON(job)
}
