package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Org      string `json:"org"`
	Location string `json:"loc"`
}

func main() {

	resp, err := http.Get("http://ipinfo.io/json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var ipInfo IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("IP:", ipInfo.IP)
	fmt.Println("City:", ipInfo.City)
	fmt.Println("Country:", ipInfo.Country)
	fmt.Println("Organization:", ipInfo.Org)
	fmt.Println("Location:", ipInfo.Location)
}
