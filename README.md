# 🌍 Operational Meteorology in CI/CD  

**A GitHub Actions Workflow for Real-Time Weather Intelligence**  
 Bridges meteorological data with continuous integration and deployment pipelines.

This project, **Global Weather Fetcher** implements a modular Go-based automation pipeline that retrieves real-time weather data across global cities, formats it for human readability, and logs results in timestamped `.txt` and `.json` files. Integrated with GitHub Actions, it supports scheduled execution, CI validation, and artifact retention—demonstrating fluency in cloud-native tooling, DevOps automation, and systems-informed design.


**Tech Stack:** Go (Golang), GitHub Actions, OpenWeatherMap API, JSON, Cron, CI/CD

**Focus Areas:** Automation, Observability, Artifact Management, Resilience Engineering


💡 **Problem Statement**

In DevOps environments, real-time environmental data can inform deployment strategies, alerting systems, and infrastructure scaling decisions. This project addresses the need for a lightweight, portable, and automatable solution to fetch and log weather data across diverse geographies—without relying on heavyweight monitoring stacks.

🧠 **Solution Design**

•	**Go CLI Tool:** Developed and architected a resilient Go script (city.go) that interfaces with the OpenWeatherMap API, 
    parses JSON responses, and formats output with timezone-aware sunrise/sunset conversion.

•	**Logging Engine:** Implemented structured logging to both .txt and .json formats, with timestamped filenames and organized 
    directory structure for historical tracking.

•	**GitHub Actions Workflow:** Built a CI/CD pipeline that triggers on push, hourly schedule, or manual dispatch. The workflow 
    runs the Go script, validates output, and uploads logs as downloadable artifacts.

•	**Error Handling:** Integrated graceful error detection for malformed API responses, missing keys, and unsupported city      
      formats—ensuring fault tolerance across runs.



### 🧠 Key Features

1. **Multi-City Coverage:** Fetches weather for London, Tokyo, Lagos, Sydney, and more 

2. **Timestamped Logging:** Saves logs in `logs/weather_log_<timestamp>.txt` and `.json`    

3. **GitHub Actions Integration:** CI/CD pipeline with push, schedule, and manual triggers   

4. **API Resilience:** Handles 400/401 errors and unsupported city formats gracefully    

5. **Timezone-Aware Conversion:** Converts sunrise/sunset from UTC to local time using API-provided offsets 

6. **Artifact Management:** Logs retained as downloadable artifacts for audit and analysis             


### 📈 Impact & Extensibility
    
    This project demonstrates strategic alignment with DevOps, automation, and transformation goals

•	**Observability:** Enables lightweight monitoring of environmental conditions across deployment zones.

•	**Automation:** Reduces manual overhead in data collection and formatting.

•	**Scalability:** Easily extendable to include forecasts, alerts, or dashboard integrations.



### Artefactory Output:

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


### 🔮 Next Steps

These next steps signal a strategic evolution of your weather-fetching pipeline from a standalone automation tool into a **production-grade observability and deployment system**. The following are the envisioned implications within a DevOps context:


### 📣 Integrate Slack/Email Alerting for Threshold-Based Weather Events

**Implications:**

- **Event-Driven Automation**: Transforms passive data collection into active response. For example, triggering alerts when 
  temperatures exceed thresholds or storms are detected.

- **Incident Readiness**: Aligns with SRE practices by enabling proactive notifications that can inform infrastructure scaling, 
failover decisions, or user-facing service adjustments.

- **Team Collaboration**: Embeds environmental intelligence directly into communication channels, improving situational awareness 
  across engineering, ops, and product teams.


### ☁️ Push Logs to Cloud Storage (Azure Blob, AWS S3)

**Implications:**

- **Durable Retention**: Moves from ephemeral GitHub artifacts to persistent, scalable storage—critical for compliance, 
  auditability, and long-term analysis.

- **Cross-System Integration**: Enables downstream consumption by data lakes, analytics platforms, or ML pipelines.

- **Cost-Efficient Archiving**: Supports tiered storage strategies (e.g., hot vs. cold data) for optimized resource usage in 
  production environments.


### 📊 Build a Dashboard for Visualizing Historical Weather Trends

**Implications:**

- **Observability Layer**: Introduces a visual interface for tracking environmental metrics over time, aligning with DevOps 
  pillars of monitoring and feedback.

- **Decision Support**: Empowers teams to correlate weather patterns with system performance, user behavior, or deployment 
  outcomes.

- **Stakeholder Engagement**: Makes data accessible to non-technical audiences (e.g., product managers, operations leads),  
  fostering cross-functional insight.


### 📦 Containerize the Service for Deployment in Edge or Hybrid Environments

**Implications:**

- **Portability & Scalability**: Enables consistent deployment across cloud, on-prem, and edge nodes—ideal for distributed 
  systems or latency-sensitive use cases.

- **Microservice Readiness**: Prepares the weather-fetcher for integration into larger service meshes or orchestrated 
  environments (e.g., Kubernetes).

- **Resilience & Isolation**: Improves fault tolerance and resource control, especially when deployed alongside other 
  observability or automation agents.


### 🧠 Strategic Summary

Together, these steps elevate the project from a developer utility to a **DevOps-aligned telemetry and alerting system**, which are a reflection of maturity in:

- **Systems Thinking**: Connecting environmental data to operational decisions  
- **Infrastructure Awareness**: Designing for scale, resilience, and modularity  
- **Transformation Strategy**: Embedding observability and automation into real-world workflows

