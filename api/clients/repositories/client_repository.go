package repositories

import (
	ClientModels "api-rest/api/clients/models"
	"api-rest/config/database"
)

func CreateClientRepository(client *ClientModels.Client) (uint, error) {
	db := database.ConectDb()
	clientResult := db.Create(&client)

	if clientResult.Error != nil {
		panic(clientResult.Error)
	}

	return client.ID, nil
}
