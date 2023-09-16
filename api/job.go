package api

import (
	"time"

	"github.com/esalavat/gojobsapi/models"
)

type Job struct {
	Id          int       `json:"id"`
	Company     string    `json:"company"`
	JobTitle    string    `json:"jobTitle"`
	DateApplied time.Time `json:"dateApplied"`
	JobUrl      string    `json:"jobUrl"`
}

func CreateResponseJob(jobModel models.Job) Job {
	return Job{
		Id:          jobModel.Id,
		Company:     jobModel.Company,
		JobTitle:    jobModel.JobTitle,
		DateApplied: jobModel.DateApplied,
		JobUrl:      jobModel.JobUrl,
	}
}

type UpdateJob struct {
	Company     string    `json:"company"`
	JobTitle    string    `json:"jobTitle"`
	DateApplied time.Time `json:"dateApplied"`
	JobUrl      string    `json:"jobUrl"`
}
