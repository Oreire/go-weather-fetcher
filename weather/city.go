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

func getWeather(city string, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}
	return &weather, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Missing 'city' parameter", http.StatusBadRequest)
		return
	}

	weather, err := getWeather(city, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf(`
		<h2>📍 %s, %s</h2>
		<p>🌤️ Condition: %s (%s)</p>
		<p>🌡️ Temperature: %.2f°C (Feels like %.2f°C)</p>
		<p>📉 Min/Max: %.2f°C / %.2f°C</p>
		<p>💧 Humidity: %d%% | Pressure: %dhPa</p>
		<p>🌬️ Wind: %.2fm/s from %d°</p>
		<p>🌅 Sunrise: %s | 🌇 Sunset: %s</p>
	`,
		weather.Name, weather.Sys.Country,
		weather.Weather[0].Main, weather.Weather[0].Description,
		weather.Main.Temp, weather.Main.FeelsLike,
		weather.Main.TempMin, weather.Main.TempMax,
		weather.Main.Humidity, weather.Main.Pressure,
		weather.Wind.Speed, weather.Wind.Deg,
		time.Unix(weather.Sys.Sunrise, 0).Format("15:04"),
		time.Unix(weather.Sys.Sunset, 0).Format("15:04"),
	)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)
	fmt.Println("🌍 Server running at http://localhost:8080/weather?city=London")
	http.ListenAndServe(":8080", nil)
}

