package transformer

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const currentWeatherJSON = `
{"coord": { "lon": 139,"lat": 35},
  "weather": [
    {
      "id": 800,
      "main": "Clear",
      "description": "clear sky",
      "icon": "01n"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 281.52,
    "feels_like": 278.99,
    "temp_min": 280.15,
    "temp_max": 283.71,
    "pressure": 1016,
    "humidity": 93
  },
  "wind": {
    "speed": 0.47,
    "deg": 107.538
  },
  "clouds": {
    "all": 2
  },
  "dt": 1560350192,
  "sys": {
    "type": 3,
    "id": 2019346,
    "message": 0.0065,
    "country": "JP",
    "sunrise": 1560281377,
    "sunset": 1560333478
  },
  "timezone": 32400,
  "id": 1851632,
  "name": "Shuzenji",
  "cod": 200
}
`

const forecastJSON = `
{
  "city": {
    "id": 2643743,
    "name": "London",
    "coord": {
      "lon": -0.1258,
      "lat": 51.5085
    },
    "country": "GB",
    "timezone": 3600
  },
  "cod": "200",
  "message": 0.7809187,
  "cnt": 7,
  "list": [
    {
      "dt": 1568977200,
      "sunrise": 1568958164,
      "sunset": 1569002733,
      "temp": {
        "day": 293.79,
        "min": 288.85,
        "max": 294.47,
        "night": 288.85,
        "eve": 290.44,
        "morn": 293.79
      },
      "feels_like": {
        "day": 278.87,
        "night": 282.73,
        "eve": 281.92,
        "morn": 278.87
      },
      "pressure": 1025.04,
      "humidity": 42,
      "weather": [
        {
          "id": 800,
          "main": "Clear",
          "description": "sky is clear",
          "icon": "01d"
        }
      ],
      "speed": 4.66,
      "deg": 102,
      "clouds": 0
    }
  ]
}`

func TestCurrentWeatherTransformer(t *testing.T) {
	expected := Weather{
		City:        "Shuzenji",
		CurrentTemp: 281.52,
		Rng: tempRange{
			Min: 280.15,
			Max: 283.71,
		},
	}

	var mapped map[string]interface{}
	json.Unmarshal([]byte(currentWeatherJSON), &mapped)

	actual := *ToWeather(mapped)

	assert.Equal(t, expected, actual)
}

func TestCurrentWeatherInvalidResponse(t *testing.T) {
	invalidJSON := `{"name": "Shuzenji"}`

	var mapped map[string]interface{}
	json.Unmarshal([]byte(invalidJSON), &mapped)

	assert.Panics(t, func() {
		ToWeather(mapped)
	}, mapped)
}

func TestForecastTransformer(t *testing.T) {
	expected := Forecast{
		City: "London",
		Temps: []temperature{
			{
				Date:  time.Unix(int64(1568977200), 0),
				Day:   293.79,
				Night: 288.85,
				Eve:   290.44,
				Morn:  293.79,
				Rng:   tempRange{Min: 288.85, Max: 294.47},
			},
		},
	}

	var mapped map[string]interface{}
	json.Unmarshal([]byte(forecastJSON), &mapped)

	actual := *ToForecast(mapped)

	assert.Equal(t, expected, actual)
}

func TestForecastInvalidResponse(t *testing.T) {
	invalidJSON := `{"aaa": "Message"}`

	var mapped map[string]interface{}
	json.Unmarshal([]byte(invalidJSON), &mapped)

	assert.Panics(t, func() {
		ToForecast(mapped)
	}, mapped)
}
