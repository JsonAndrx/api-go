package location

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetLocation(ip string) (string, error) {
	ipapiClient := http.Client{}
	fmt.Println(ip, "yeison")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/country_name/", ip), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := ipapiClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return (string(body)), nil
}
