package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	http.HandleFunc("/carbon", carbonHandler)
	http.HandleFunc("/my_response", myResponse)
	http.HandleFunc("/health", healthCheck)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func fetch_carbon_intensity() (CarbonIntensityResponse, error) {
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	resp, err := http.Get("https://api.carbonintensity.org.uk/regional/postcode/TF8")

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
		return CarbonIntensityResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var data CarbonIntensityResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal("Error while fetching carbon intensity data: %s", err)
	}

	return data, nil
}

type Response struct {
	Message string `json:"message"`
}

func carbonHandler(w http.ResponseWriter, r *http.Request) {
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	resp, err := fetch_carbon_intensity()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Error while fetching API data")
		res_msg := fmt.Sprintf("Error while fetching API data: %s", err)
		http.Error(w, res_msg, 400)
		return
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
