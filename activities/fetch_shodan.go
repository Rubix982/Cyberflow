package activities

import (
	"context"
	"log"
)

type ShodanResponse struct {
	IP        string   `json:"ip"`
	Hostnames []string `json:"hostnames"`
	Org       string   `json:"org"`
	ISP       string   `json:"isp"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
}

func FetchShodan(ctx context.Context) (interface{}, error) {

	log.Println("Fetching data from Shodan...")

	return map[string]string{"result": "Shodan stub"}, nil
}
