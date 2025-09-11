package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getWeather(city string, apiKey string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Printf("🌤️ %s: %.2f°C\n", weather.Name, weather.Main.Temp)
}

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	cities := []string{"London", "Tokyo", "New York", "Lagos", "Sydney"}

	for _, city := range cities {
		getWeather(city, apiKey)
	}
}

// To run this code, ensure you have set the environment variable OPENWEATHER_API_KEY with your OpenWeatherMap API key.