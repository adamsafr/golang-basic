package router

import (
	"net/http"

	"github.com/adamsafr/golang-basic/pkg/http/handler"
)

// Route ...
type Route struct {
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// GetRouteList ...
func GetRouteList() []Route {
	return []Route{
		{"/current-weather", "GET", handler.CurrentWeatherHandler},
		{"/daily-forecast", "GET", handler.DailyForecastHandler},
	}
}
