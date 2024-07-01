package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ShellyAPIClient struct {
	ApiKey string
}

func (s ShellyAPIClient) set_api_key(key string) {
	s.ApiKey = key
}

func (s ShellyAPIClient) fetch(id string) ShellyResponse {
	id = "d48afc400484"
	req := ShellyRequest{s.ApiKey, id}
	json_req, err := json.Marshal(req)
	resp, err := http.Post("shelly-100-eu.shelly.cloud/device/status", "application/json", bytes.NewBuffer(json_req))

	if err != nil {
		fmt.Printf("Error while fetching Shelly data: %s", err)
	}

	defer resp.Body.Close()

	var data ShellyResponse

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		fmt.Printf("Error while fetching carbon intensity data: %s", err)
	}

	return data
}

type ShellyRequest struct {
	ApiKey string `json: "auth_key"`
	Id     string `json: "id"`
}

type ShellyResponse struct {
	IsOk bool       `json:"is_ok"`
	Data ShellyData `json:"data"`
}

type ShellyData struct {
	Online bool `json:"online"`
}

type DeviceStatus struct {
	Ts      float32     `json:"ts"`
	Cloud   CloudStatus `json:"cloud"`
	Wifi    Wifi        `json:"wifi"`
	Code    string      `json:"code"`
	Switch0 Switch      `json:"switch:0"`
}

type CloudStatus struct {
	IsConnected bool `json:"connected"`
}

type Wifi struct {
	StaticIp string `json:"sta_ip"`
	Status   string `json:"status"`
	Ssid     string `json:"ssid"`
	rssi     int    `json:"rssi"`
}

type Switch struct {
	Apower float32 `json:"apower"`
}
