# 🌍 Operational Meteorology in CI/CD  
**A GitHub Actions Workflow for Real-Time Weather Intelligence**  
Bridges meteorological data with continuous integration and deployment pipelines.

This project, **Global Weather Fetcher** implements a modular Go-based automation pipeline that retrieves real-time weather data across global cities, formats it for human readability, and logs results in timestamped `.txt` and `.json` files. Integrated with GitHub Actions, it supports scheduled execution, CI validation, and artifact retention—demonstrating fluency in cloud-native tooling, DevOps automation, and systems-informed design.


### 🔧 Role & Responsibilities
- Architected a resilient CLI tool in Go to interface with OpenWeatherMap API
- Implemented structured logging and timezone-aware formatting
- Built a CI/CD pipeline using GitHub Actions with cron scheduling and artifact uploads
- Designed fault-tolerant error handling for API failures and malformed responses

### 🧠 Technical Highlights
 ______________________________________________________________________________________________________________
| Capability                     | Description                                                                 |
|-------------------------------|----------------------------------------------------------------------------- |
| 🌐 Multi-City Coverage         | Fetches weather for London, Tokyo, Lagos, Sydney, and more                 |
| 📦 Timestamped Logging         | Saves logs in `logs/weather_log_<timestamp>.txt` and `.json`               |
| 🔁 GitHub Actions Integration  | CI/CD pipeline with push, schedule, and manual triggers                    |
| 🧪 API Resilience              | Handles 400/401 errors and unsupported city formats gracefully             |
| 🌅 Timezone-Aware Conversion   | Converts sunrise/sunset from UTC to local time using API-provided offsets |
| 📁 Artifact Management         | Logs retained as downloadable artifacts for audit and analysis             |


### 📈 Impact & Extensibility

- Enables lightweight observability across deployment zones
- Reduces manual overhead in environmental data collection
- Easily extendable to include forecasts, alerts, or dashboard integrations
- Demonstrates strategic alignment with DevOps, automation, and transformation goals


Artefactory Output:

https://github.com/Oreire/go-prog/actions/runs/17631236409/artifacts/3981369294

📍 London, GB
🌤️ Condition: Clouds (few clouds)
🌡️ Temperature: 13.09°C (Feels like 12.78°C)
📉 Min/Max: 11.62°C / 13.86°C
💧 Humidity: 89% | Pressure: 1001hPa
🌬️ Wind: 6.17m/s from 240°
🌅 Sunrise: 06:29 | 🌇 Sunset: 19:25

📍 Austria, AT
🌤️ Condition: Clouds (overcast clouds)
🌡️ Temperature: 10.74°C (Feels like 10.30°C)
📉 Min/Max: 10.74°C / 10.74°C
💧 Humidity: 93% | Pressure: 1011hPa
🌬️ Wind: 0.63m/s from 176°
🌅 Sunrise: 06:38 | 🌇 Sunset: 19:28

📍 Tokyo, JP
🌤️ Condition: Clouds (broken clouds)
🌡️ Temperature: 31.27°C (Feels like 36.52°C)
📉 Min/Max: 29.75°C / 32.02°C
💧 Humidity: 64% | Pressure: 1014hPa
🌬️ Wind: 6.69m/s from 180°
🌅 Sunrise: 05:20 | 🌇 Sunset: 17:55

⚠️ API error for New York City: 400 Bad Request

📍 Lagos, NG
🌤️ Condition: Clouds (overcast clouds)
🌡️ Temperature: 23.71°C (Feels like 24.59°C)
📉 Min/Max: 23.71°C / 23.71°C
💧 Humidity: 94% | Pressure: 1014hPa
🌬️ Wind: 2.18m/s from 226°
🌅 Sunrise: 06:36 | 🌇 Sunset: 18:47

📍 Sydney, AU
🌤️ Condition: Rain (light intensity shower rain)
🌡️ Temperature: 12.08°C (Feels like 11.82°C)
📉 Min/Max: 11.32°C / 14.38°C
💧 Humidity: 95% | Pressure: 1011hPa
🌬️ Wind: 12.35m/s from 220°
🌅 Sunrise: 06:00 | 🌇 Sunset: 17:43


