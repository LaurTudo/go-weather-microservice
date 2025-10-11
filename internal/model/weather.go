package model

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float32 `json:"temp_c"`
		UpdatedAt string  `json:"last_updated"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Days []struct {
			Hours []struct {
				TimeEpoch int     `json:"time_epoch"`
				TempC     float32 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
			Day struct {
				MaxTempC        float32 `json:"maxtemp_c"`
				DailyChanceRain int     `json:"daily_chance_of_rain"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
