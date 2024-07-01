package main

import (
	"time"
)

type CarbonIntensityResponse struct {
	RegionId  int    `json:"regionid"`
	Dnoregion string `json:"dnoregion"`
	Shortname string `json:"shortname"`
	Postcode  string `json:"postcode"`
	Data      []Data `json:data`
}

type Data struct {
	From          time.Time       `json:"from"`
	To            time.Time       `json:"to"`
	Intensity     Intensity       `json:"intensity"`
	GenerationMix []GenerationMix `json:"generationmix"`
}

type Intensity struct {
	Forecast int    `json:"forecast"`
	Index    string `json:"index"`
}

type GenerationMix struct {
	Fuel string  `json:"fuel"`
	Perc float32 `json:"perc"`
}

type MyResponse struct {
	Name            string       `json:"name" fake:"{firstname}"`
	TimeTaken       int          `json:"time_taken" fake:"{number}"`
	ElectricityUsed string       `json:"electricity_used" fake:"{number}"`
	Cost            float32      `json:"cost" fake:"{number}"`
	CarbonIntensity string       `json:"carbon_intensity" fake:"{number}"`
	MemoryUsed      string       `json:"memory_used" fake:"{number}"`
	TimeSeries      []TimeSeries `json:"timeseries" fakesize:"2,10"`
}

type TimeSeries struct {
	Timestamp       time.Time `json:"timestamp"`
	ElectricityUsed string    `json:"electricity_used" fake:"{number}"`
	cost            float32   `json:"cost" fake:"{currency}"`
	CarbonIntensity string    `json:"carbon_intensity" fake:"{number}"`
	MemoryUsed      string    `json:"memory_used" fake:"{number}"`
}
