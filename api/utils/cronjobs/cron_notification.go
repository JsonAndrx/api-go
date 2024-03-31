package cronjobs

import (
	"api-rest/api/clients/repositories"
	"fmt"

	"github.com/robfig/cron/v3"
	"api-rest/api/utils/notifications"
)

func StartCronNotification() {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		clients, err := repositories.GetClientsExpireDate()
		fmt.Println(len(clients), "clients expire in the next minute")
		if err != nil {
			panic(err)
		}

		for _, client := range clients {
			notifications.SendMessage(client)
		}
	})
	c.Start()
}
