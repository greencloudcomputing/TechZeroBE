package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	http.Handle("/carbon", enableCORS(http.HandlerFunc(carbonHandler)))
	http.Handle("/my_response", enableCORS(http.HandlerFunc(myResponse)))
	http.Handle("/health", enableCORS(http.HandlerFunc(healthCheck)))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func fetch_carbon_intensity() (CarbonIntensityResponse, error) {
	fmt.Printf("Received response!")
	resp, err := http.Get("https://api.carbonintensity.org.uk/regional/postcode/CV8")

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
		return CarbonIntensityResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
	}

	var data CarbonIntensityResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
	}

	return data, nil
}

type Response struct {
	Message string `json:"message"`
}

func carbonHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := fetch_carbon_intensity()

	if err != nil {
		fmt.Printf("Error while fetching API data: %s", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	}
}

func myResponse(w http.ResponseWriter, r *http.Request) {
	var f MyResponse
	err := gofakeit.Struct(&f)

	if err != nil {
		fmt.Printf("Error while generating my response: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f)
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
