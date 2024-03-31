package notifications

import (
	"fmt"
	"os"

	"api-rest/api/clients/models"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage(clientData models.Client) {
	fmt.Println("Enviando mensaje a: ", clientData)
	os.Setenv("TWILIO_ACCOUNT_SID", "AC2917c7ec5db82e70ba3bfc0de6df4513")
	os.Setenv("TWILIO_AUTH_TOKEN", "442e6e282c5f555a0c1e91ea0f209e86")

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody(fmt.Sprintf("Hola %s %s, %s", clientData.FirtsName, clientData.LastName, clientData.Description))
	params.SetTo("whatsapp:+" + clientData.Phone)

	fmt.Println(params)
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		panic(err)
	} else {
		if resp.Sid != nil {
			fmt.Println("Mensaje enviado con éxito, ID:", *resp.Sid)
		} else {
			fmt.Println("Mensaje enviado con éxito, ID:", resp.Sid)
		}
	}
}
