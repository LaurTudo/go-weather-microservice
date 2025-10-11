package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"goweathermicroservice/internal/model"
)

func FetchWeather(apiKey, city string) (*model.WeatherResponse, error) {
	if apiKey == "" {
		return nil, errors.New("missing API key (set WEATHER_API_KEY env variable)")
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var errResp map[string]any
		json.Unmarshal(body, &errResp)

		if apiErr, ok := errResp["error"].(map[string]any); ok {
			return nil, fmt.Errorf("API error: %v", apiErr["message"])
		}
		return nil, fmt.Errorf("unexpected API response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var data model.WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &data, nil
}
