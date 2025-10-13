package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"goweathermicroservice/internal/api"
	"goweathermicroservice/internal/model"
	"goweathermicroservice/server"

	"github.com/joho/godotenv"
)

func Run() {
	// load api key from env file
	if err := godotenv.Load("secrets.env"); err != nil {
		fmt.Println(".env file could not be loaded:", err)
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set WEATHER_API_KEY environment variable.")
		os.Exit(1)
	}

	// CLI flags
	mode := flag.String("mode", "cli", "Run mode: cli or server")
	city := flag.String("city", "London", "City name for weather forecast")
	flag.Parse()

	// if flag mode server, call RunServer func
	if *mode == "server" {
		server.RunServer(apiKey)
		return
	}

	// fetch weather
	weather, err := api.FetchWeather(apiKey, *city)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	printForecast(weather)
}

func printForecast(w *model.WeatherResponse) {
	fmt.Printf("\n%s, %s\n", w.Location.Name, w.Location.Country)
	fmt.Printf("Current: %.1f°C (%s)\n", w.Current.TempC, w.Current.Condition.Text)

	now := time.Now().Hour()
	fmt.Println("\nForecast for the rest of the day:")
	for _, hour := range w.Forecast.Days[0].Hours {
		h := time.Unix(int64(hour.TimeEpoch), 0)
		if h.Hour() >= now {

			fmt.Printf("  %02d:00 -> %.1f°C (%s)\n", h.Hour(), hour.TempC, hour.Condition.Text)

		}
	}
	fmt.Println()
}
