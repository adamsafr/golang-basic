package transformer

import "time"

type tempRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// Weather ...
type Weather struct {
	City        string    `json:"city"`
	CurrentTemp float64   `json:"currentTemp"`
	Rng         tempRange `json:"range"`
}

type temperature struct {
	Date  time.Time `json:"date"`
	Day   float64   `json:"day"`
	Night float64   `json:"night"`
	Eve   float64   `json:"eve"`
	Morn  float64   `json:"morn"`
	Rng   tempRange `json:"range"`
}

// Forecast ...
type Forecast struct {
	City  string        `json:"city"`
	Temps []temperature `json:"temps"`
}

// ToWeather ...
func ToWeather(mapped map[string]interface{}) *Weather {
	res := mapped["main"].(map[string]interface{})

	return &Weather{
		City:        mapped["name"].(string),
		CurrentTemp: res["temp"].(float64),
		Rng: tempRange{
			Min: res["temp_min"].(float64),
			Max: res["temp_max"].(float64),
		},
	}
}

// ToForecast ...
func ToForecast(mapped map[string]interface{}) *Forecast {
	var temps []temperature

	for _, val := range mapped["list"].([]interface{}) {
		item := val.(map[string]interface{})
		temp := item["temp"].(map[string]interface{})

		temps = append(temps, temperature{
			Date:  time.Unix(int64(item["dt"].(float64)), 0),
			Day:   temp["day"].(float64),
			Night: temp["night"].(float64),
			Eve:   temp["eve"].(float64),
			Morn:  temp["morn"].(float64),
			Rng: tempRange{
				Min: temp["min"].(float64),
				Max: temp["max"].(float64),
			},
		})
	}

	city := mapped["city"].(map[string]interface{})["name"].(string)

	return &Forecast{City: city, Temps: temps}
}
