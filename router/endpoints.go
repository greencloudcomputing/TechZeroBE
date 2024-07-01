package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"main/external"
	"main/models"
)

func myResponse(w http.ResponseWriter, r *http.Request) {
	var f models.MyResponse
	err := gofakeit.Struct(&f)

	if err != nil {
		fmt.Printf("Error while generating my response: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f)
}

type Response struct {
	Message string `json:"message"`
}

func EnergyHandler(w http.ResponseWriter, r *http.Request) {
	query_params := r.URL.Query()
	from_date, to_date, function_id := query_params.Get("from"), query_params.Get("to"), query_params.Get("functionId")

	data := external.GetCarbonIntensityApi(from_date, to_date, function_id)

	meme := external.CalcCarbonIntensity(data.Data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meme)
}

type MyResp struct {
	TotalCarbonIntensity int `json:"total_carbon_intensity"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!"))
}
