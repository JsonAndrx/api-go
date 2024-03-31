package repositories

import (
	ClientModels "api-rest/api/clients/models"
	"api-rest/config/database"
	"time"
)

func CreateClientRepository(client *ClientModels.Client) (uint, error) {
	db := database.ConectDb()
	clientResult := db.Create(&client)

	if clientResult.Error != nil {
		panic(clientResult.Error)
	}

	return client.ID, nil
}

func GetClientsExpireDate() ([]ClientModels.Client, error) {
	db := database.ConectDb()
	clients := []ClientModels.Client{}
	now := time.Now()
	previousMinute := now.Add(time.Minute * -1)

	db.Where("date >= ? AND date <= ?", previousMinute, now).Find(&clients)

	return clients, nil
}
