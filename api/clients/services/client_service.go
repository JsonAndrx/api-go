package services

import (
	"fmt"
	"net/http"
	"path/filepath"

	"strings"
	"time"

	// "api-rest/api/clients/models"
	// "api-rest/api/clients/repositories"
	"api-rest/api/clients/types"
	"api-rest/api/utils/response"
	"api-rest/config/database"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func UploadClientSerive(c *gin.Context) {
	fileHeader, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file is uploaded"})
		return
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".xlsx" && ext != ".xls" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("The file must be in xlsx or xls format", ""))
		return
	}

	path := filepath.Join("files/", fileHeader.Filename)
	if err := c.SaveUploadedFile(fileHeader, path); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error file", err.Error()))
		return
	}

	excelFile, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := excelFile.GetRows("Hoja1")
	clients := []types.ClientCreateRequest{}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		clientsType := types.ClientCreateRequest{
			UsuarioID:   1,
			Nombre:      row[0],
			Apellido:    row[1],
			Telefono:    row[2],
			Fecha:       row[3],
			Hora:		row[4],
			Descripcion: row[5],
		}

		fields := map[string]string{
			"Nombre":      clientsType.Nombre,
			"Apellido":    clientsType.Apellido,
			"Telefono":    clientsType.Telefono,
			"Fecha":       clientsType.Fecha,
			"Hora":        clientsType.Hora,
			"Descripcion": clientsType.Descripcion,
		}

		for field, value := range fields {
			if value == "" {
				c.JSON(http.StatusBadRequest, response.ErrorResponse("Some value of the "+field+" field is empty", ""))
				return
			}
		}

		clients = append(clients, clientsType)
	}

	responseCreate, succesClients := CreateClients(clients)
	if !succesClients {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(responseCreate, ""))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(responseCreate, ""))
}

func CreateClients(clients []types.ClientCreateRequest) (string, bool) {
	var valueStrings []string
	var valueArgs []interface{}

	db := database.ConectDb()

	if len(clients) == 0 {
		return "Not exist clients to create", false
	} else if len(clients) > 500 {
		return "The limit of upload clients to create is 500", false
	}

	for _, client := range clients {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")

		layout := "01-02-06 15:04"
		t, err := time.Parse(layout, client.Fecha + " " + client.Hora)
		if err != nil {
			panic(err)
		}

		valueArgs = append(valueArgs, client.UsuarioID, client.Nombre, client.Apellido, client.Telefono, t, client.Descripcion)
	}

	stmt := fmt.Sprintf("INSERT INTO clients (user_id, firts_name, last_name, phone, date, description) VALUES %s", strings.Join(valueStrings, ","))

	err := db.Exec(stmt, valueArgs...).Error
	if err != nil {
		return "Error to create the clients", false
	}

	return fmt.Sprintf("Se crearon %d clientes", len(clients)), true
}
