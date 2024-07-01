package external

import (
	"encoding/json"
	"fmt"

	"main/models"
	"net/http"
)

func GetCarbonIntensityApi(from string, to string, function_id string) models.CarbonResponse {
	url := fmt.Sprintf("https://api.carbonintensity.org.uk/regional/intensity/%s/%s/postcode/TF8", from, to)
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
	}

	fmt.Println(resp.Status)

	defer resp.Body.Close()

	var data models.CarbonResponse

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
	}

	return data
}
