package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adamsafr/golang-basic/pkg/http/response"
	"github.com/adamsafr/golang-basic/pkg/http/transformer"
	"github.com/adamsafr/golang-basic/pkg/service/openweather"
	"github.com/adamsafr/golang-basic/pkg/util/str"
)

// CurrentWeatherHandler ...
func CurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	res := response.Create(w)
	q := r.URL.Query().Get("city")

	if str.IsEmpty(q) {
		res.BadRequest("City cannot be empty.")
		return
	}

	data := openweather.CreateProvider().GetCurrent(q)

	var mapped map[string]interface{}
	err := json.Unmarshal(data, &mapped)

	if err != nil {
		res.Error("Cannot decode json string", http.StatusInternalServerError)
		return
	}

	res.ToJSON(transformer.ToWeather(mapped))
}

// DailyForecastHandler ...
func DailyForecastHandler(w http.ResponseWriter, r *http.Request) {
	res := response.Create(w)
	q := r.URL.Query().Get("city")

	if str.IsEmpty(q) {
		res.BadRequest("City cannot be empty.")
		return
	}

	data := openweather.CreateProvider().GetDaily(q, 7)

	var mapped map[string]interface{}
	err := json.Unmarshal(data, &mapped)

	if err != nil {
		res.Error("Cannot decode json string", http.StatusInternalServerError)
		return
	}

	res.ToJSON(transformer.ToForecast(mapped))
}
