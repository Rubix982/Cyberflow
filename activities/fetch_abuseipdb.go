package activities

import (
	"context"
	"log"
)

func FetchAbuseIPDB(ctx context.Context) (interface{}, error) {
	log.Println("Fetching data from AbuseIPDB...")

	return map[string]string{"result": "AbuseIPDB stub"}, nil
}
