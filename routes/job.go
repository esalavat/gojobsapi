package routes

import (
	"log"

	"github.com/esalavat/gojobsapi/api"
	"github.com/esalavat/gojobsapi/db"
	"github.com/esalavat/gojobsapi/models"
	"github.com/gofiber/fiber/v2"
)

func CreateJob(c *fiber.Ctx) error {
	var job api.Job

	if err := c.BodyParser(&job); err != nil {
		log.Println("Error on parse");
		log.Println(err);
		return c.Status(400).JSON(err.Error())
	}

	var modelJob = api.CreateModelJob(job);

	db.Database.Create(&modelJob)

	responseJob := api.CreateResponseJob(modelJob)

	return c.Status(200).JSON(responseJob)
}

func GetJobs(c *fiber.Ctx) error {
	jobs := []models.Job{}

	db.Database.Model(&models.Job{}).Preload("Updates").Find(&jobs)

	responseJobs := []api.Job{}
	for _, job := range jobs {
		responseJob := api.CreateResponseJob(job)
		responseJobs = append(responseJobs, responseJob)
	}

	return c.Status(200).JSON(responseJobs)
}

func GetJob(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Could not parse :id param.")
	}

	var job models.Job

	db.Database.Model(&models.Job{}).Preload("Updates").Find(&job, "id = ?", id)

	if job.Id == 0 {
		return c.Status(400).JSON("User does not exist.")
	}

	responseJob := api.CreateResponseJob(job)

	return c.Status(200).JSON(responseJob)
}

func UpdateJob(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Could not parse :id param.")
	}

	var job models.Job

	db.Database.Model(&models.Job{}).Preload("Updates").Find(&job, "id = ?", id)

	if job.Id == 0 {
		return c.Status(400).JSON("User does not exist.")
	}

	var updateData api.UpdateJob

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	job.Company = updateData.Company
	job.JobTitle = updateData.JobTitle
	job.DateApplied = updateData.DateApplied
	job.JobUrl = updateData.JobUrl
	job.Updates = api.CreateModelUpdates(updateData.Updates)

	db.Database.Save(&job)

	responseJob := api.CreateResponseJob(job)

	return c.Status(200).JSON(responseJob)
}

func DeleteJob(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Could not parse :id param.")
	}

	var job models.Job

	db.Database.Model(&models.Job{}).Preload("Updates").Find(&job, "id = ?", id)

	if job.Id == 0 {
		return c.Status(400).JSON("User does not exist.")
	}

	var ids []int
	for _, update := range job.Updates {
		ids = append(ids, update.Id)
	}

	db.Database.Delete(&models.Update{}, ids)
	db.Database.Delete(job)

	return c.Status(204).Send(nil)
}
