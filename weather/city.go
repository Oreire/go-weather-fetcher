package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type WeatherResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
}

func getWeather(city string, apiKey string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ Error fetching weather for %s: %v\n", city, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("⚠️ API error for %s: %s\n", city, resp.Status)
		return
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Printf("🛑 Decode error for %s: %v\n", city, err)
		return
	}

	fmt.Printf("\n📍 %s, %s\n", weather.Name, weather.Sys.Country)
	fmt.Printf("🌤️ Condition: %s (%s)\n", weather.Weather[0].Main, weather.Weather[0].Description)
	fmt.Printf("🌡️ Temperature: %.2f°C (Feels like %.2f°C)\n", weather.Main.Temp, weather.Main.FeelsLike)
	fmt.Printf("📉 Min/Max: %.2f°C / %.2f°C\n", weather.Main.TempMin, weather.Main.TempMax)
	fmt.Printf("💧 Humidity: %d%% | Pressure: %dhPa\n", weather.Main.Humidity, weather.Main.Pressure)
	fmt.Printf("🌬️ Wind: %.2fm/s from %d°\n", weather.Wind.Speed, weather.Wind.Deg)
	fmt.Printf("🌅 Sunrise: %s | 🌇 Sunset: %s\n",
		time.Unix(weather.Sys.Sunrise, 0).Format("15:04"),
		time.Unix(weather.Sys.Sunset, 0).Format("15:04"))
}

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("❗ OPENWEATHER_API_KEY environment variable not set.")
		return
	}

	cities := []string{"London", "Tokyo", "New York", "Lagos", "Sydney"}

	for _, city := range cities {
		getWeather(city, apiKey)
	}
}

