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

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!"))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's a preflight request, handle it directly
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func ServeRouter() {
	http.Handle("/carbon", enableCORS(http.HandlerFunc(EnergyHandler)))
	http.Handle("/my_response", enableCORS(http.HandlerFunc(myResponse)))
	http.Handle("/health", enableCORS(http.HandlerFunc(healthCheck)))

	fmt.Print("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err.Error())
	}
}

func EnergyHandler(w http.ResponseWriter, r *http.Request) {
	query_params := r.URL.Query()
	from_date, to_date, function_id := query_params.Get("from"), query_params.Get("to"), query_params.Get("functionId")

	data := external.GetCarbonIntensityApi(from_date, to_date, function_id)

	i := 0
	for _, datapoint := range data.Data.Data {
		i += datapoint.Intensity.Forecast
	}

	myresp := MyResp{i}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myresp)
}

type MyResp struct {
	TotalCarbonIntensity int `json:"total_carbon_intensity"`
}
