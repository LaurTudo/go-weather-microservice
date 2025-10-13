package server

import (
	"encoding/json"
	"fmt"
	"goweathermicroservice/internal/api"
	"goweathermicroservice/internal/model"
	"net/http"
	"time"
)

// http server, exposing /weather
func RunServer(apiKey string) {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "missing city parameter (use something like /weather?city=Bucharest)", http.StatusBadRequest)
			return
		}

		data, err := api.FetchWeather(apiKey, city)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch weather: %v", err), http.StatusBadRequest)
			return
		}

		// WeatherSummary stucture used here, from model/weather.go
		summary := model.WeatherSummary{
			City:      data.Location.Name,
			Country:   data.Location.Country,
			TempC:     data.Current.TempC,
			Condition: data.Current.Condition.Text,
		}

		currentEpoch := time.Now().Unix()
		for _, hour := range data.Forecast.Days[0].Hours {
			h := time.Unix(int64(hour.TimeEpoch), 0)
			if int64(hour.TimeEpoch) >= currentEpoch {
				summary.Forecast = append(summary.Forecast, model.HourForecast{
					Hour:      h.Hour(),
					TempC:     hour.TempC,
					Condition: hour.Condition.Text,
				})
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(summary)
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
