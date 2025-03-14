package main

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"github.com/fatih/color"
)

// Structs to hold the API response
type WeatherResponse struct {
	Location struct {
		Name    string  `json:"name"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	} `json:"location"`

	Current struct {
		TempC      float64 `json:"temp_c"`
		Condition  struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		Humidity   int     `json:"humidity"`
		FeelsLikeC float64 `json:"feelslike_c"`
	} `json:"current"`
}

func main() {

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=a873ecc3d49f43389be172014251303&q=Dhaka")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather API is not available")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
		// Unmarshal JSON into struct
		var weatherData WeatherResponse
		err = json.Unmarshal(body, &weatherData)
		if err != nil {
			panic(err)
		}
			// Print weather details
// Define colorized output
cityColor := color.New(color.FgCyan, color.Bold)
tempColor := color.New(color.FgYellow, color.Bold)
conditionColor := color.New(color.FgGreen, color.Bold)
windColor := color.New(color.FgBlue, color.Bold)
humidityColor := color.New(color.FgMagenta, color.Bold)

// Print weather details with colors
cityColor.Printf("Weather in %s, %s:\n", weatherData.Location.Name, weatherData.Location.Country)
tempColor.Printf("Temperature: %.1f°C (Feels like %.1f°C)\n", weatherData.Current.TempC, weatherData.Current.FeelsLikeC)
conditionColor.Printf("Condition: %s\n", weatherData.Current.Condition.Text)
windColor.Printf("Wind Speed: %.1f km/h\n", weatherData.Current.WindKph)
humidityColor.Printf("Humidity: %d%%\n", weatherData.Current.Humidity)
fmt.Printf("Icon URL: https:%s\n", weatherData.Current.Condition.Icon)

}
