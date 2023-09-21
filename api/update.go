package api

import (
	"time"

	"github.com/esalavat/gojobsapi/models"
)

type Update struct {
	Id          int       `json:"id"`
	Display     string    `json:"display"`
	DateUpdated time.Time `json:"dateUpdated"`
	Notes       string    `json:"notes"`
}

func CreateResponseUpdate(updateModel models.Update) Update {
	return Update{
		Id:          updateModel.Id,
		Display:     updateModel.Display,
		DateUpdated: updateModel.DateUpdated,
		Notes:       updateModel.Notes,
	}
}

func CreateResponseUpdates(updateModels []models.Update) []Update {

	responseUpdates := []Update{}

	for _, updateModel := range updateModels {
		responseUpdate := CreateResponseUpdate(updateModel)
		responseUpdates = append(responseUpdates, responseUpdate)
	}

	return responseUpdates
}

func CreateModelUpdate(updateApi Update) models.Update {
	return models.Update{
		Id:          updateApi.Id,
		Display:     updateApi.Display,
		DateUpdated: updateApi.DateUpdated,
		Notes:       updateApi.Notes,
	}
}

func CreateModelUpdates(updatesApi []Update) []models.Update {
	modelUpdates := []models.Update{}

	for _, updateApi := range updatesApi {
		modelUpdate := CreateModelUpdate(updateApi)
		modelUpdates = append(modelUpdates, modelUpdate)
	}

	return modelUpdates
}
