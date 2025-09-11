// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"
// )

// type WeatherResponse struct {
// 	Name    string `json:"name"`
// 	Weather []struct {
// 		Main        string `json:"main"`
// 		Description string `json:"description"`
// 	} `json:"weather"`
// 	Main struct {
// 		Temp      float64 `json:"temp"`
// 		FeelsLike float64 `json:"feels_like"`
// 		TempMin   float64 `json:"temp_min"`
// 		TempMax   float64 `json:"temp_max"`
// 		Humidity  int     `json:"humidity"`
// 		Pressure  int     `json:"pressure"`
// 	} `json:"main"`
// 	Wind struct {
// 		Speed float64 `json:"speed"`
// 		Deg   int     `json:"deg"`
// 	} `json:"wind"`
// 	Sys struct {
// 		Country string `json:"country"`
// 		Sunrise int64  `json:"sunrise"`
// 		Sunset  int64  `json:"sunset"`
// 	} `json:"sys"`
// }

// func getWeather(city string, apiKey string) {
// 	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("❌ Error fetching weather for %s: %v\n", city, err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("⚠️ API error for %s: %s\n", city, resp.Status)
// 		return
// 	}

// 	var weather WeatherResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
// 		fmt.Printf("🛑 Decode error for %s: %v\n", city, err)
// 		return
// 	}

// 	fmt.Printf("\n📍 %s, %s\n", weather.Name, weather.Sys.Country)
// 	fmt.Printf("🌤️ Condition: %s (%s)\n", weather.Weather[0].Main, weather.Weather[0].Description)
// 	fmt.Printf("🌡️ Temperature: %.2f°C (Feels like %.2f°C)\n", weather.Main.Temp, weather.Main.FeelsLike)
// 	fmt.Printf("📉 Min/Max: %.2f°C / %.2f°C\n", weather.Main.TempMin, weather.Main.TempMax)
// 	fmt.Printf("💧 Humidity: %d%% | Pressure: %dhPa\n", weather.Main.Humidity, weather.Main.Pressure)
// 	fmt.Printf("🌬️ Wind: %.2fm/s from %d°\n", weather.Wind.Speed, weather.Wind.Deg)
// 	fmt.Printf("🌅 Sunrise: %s | 🌇 Sunset: %s\n",
// 		time.Unix(weather.Sys.Sunrise, 0).Format("15:04"),
// 		time.Unix(weather.Sys.Sunset, 0).Format("15:04"))
// }

// func main() {
// 	apiKey := os.Getenv("OPENWEATHER_API_KEY")
// 	if apiKey == "" {
// 		fmt.Println("❗ OPENWEATHER_API_KEY environment variable not set.")
// 		return
// 	}

// 	cities := []string{"London", "Austria", "Tokyo", "Lagos", "Sydney"}

// 	for _, city := range cities {
// 		getWeather(city, apiKey)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type WeatherResponse struct {
	Name    string `json:"name"`
	Coord   struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
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
	Timezone int64 `json:"timezone"`
}

func getWeather(city string, apiKey string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error for %s: %s", city, resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}
	return &weather, nil
}

func formatWeatherText(w *WeatherResponse) string {
	sunrise := time.Unix(w.Sys.Sunrise, 0).Add(time.Second * time.Duration(w.Timezone)).Format("15:04")
	sunset := time.Unix(w.Sys.Sunset, 0).Add(time.Second * time.Duration(w.Timezone)).Format("15:04")

	return fmt.Sprintf(
		"📍 %s, %s\n🌤️ Condition: %s (%s)\n🌡️ Temperature: %.2f°C (Feels like %.2f°C)\n📉 Min/Max: %.2f°C / %.2f°C\n💧 Humidity: %d%% | Pressure: %dhPa\n🌬️ Wind: %.2fm/s from %d°\n🌅 Sunrise: %s | 🌇 Sunset: %s\n\n",
		w.Name, w.Sys.Country,
		w.Weather[0].Main, w.Weather[0].Description,
		w.Main.Temp, w.Main.FeelsLike,
		w.Main.TempMin, w.Main.TempMax,
		w.Main.Humidity, w.Main.Pressure,
		w.Wind.Speed, w.Wind.Deg,
		sunrise, sunset,
	)
}

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("❗ OPENWEATHER_API_KEY environment variable not set.")
		return
	}

	cities := []string{"London", "Austria", "Tokyo", "New York City", "Lagos", "Sydney"}
	var logText strings.Builder
	var logJSON []WeatherResponse

	for _, city := range cities {
		weather, err := getWeather(city, apiKey)
		if err != nil {
			logText.WriteString(fmt.Sprintf("⚠️ %v\n\n", err))
			continue
		}
		logText.WriteString(formatWeatherText(weather))
		logJSON = append(logJSON, *weather)
	}

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	txtFile := fmt.Sprintf("logs/weather_log_%s.txt", timestamp)
	jsonFile := fmt.Sprintf("logs/weather_log_%s.json", timestamp)

	// Write .txt log
	if err := os.WriteFile(txtFile, []byte(logText.String()), 0644); err != nil {
		fmt.Println("❌ Error writing TXT log:", err)
	}

	// Write .json log
	jsonData, err := json.MarshalIndent(logJSON, "", "  ")
	if err != nil {
		fmt.Println("❌ Error marshaling JSON:", err)
	} else if err := os.WriteFile(jsonFile, jsonData, 0644); err != nil {
		fmt.Println("❌ Error writing JSON log:", err)
	}

	fmt.Println("✅ Weather logs saved to:")
	fmt.Println("   -", txtFile)
	fmt.Println("   -", jsonFile)
}

