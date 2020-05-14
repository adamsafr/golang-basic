package openweather

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/adamsafr/golang-basic/config"
)

// Provider ...
type Provider struct {
	Path   string
	AppKey string
}

// CreateProvider ...
func CreateProvider() *Provider {
	c := config.Get().OpenWeather

	return &Provider{
		Path:   c.Path,
		AppKey: c.AppKey,
	}
}

// GetCurrent ...
func (ow *Provider) GetCurrent(city string) []byte {
	res, err := http.Get(ow.GeneratePath("weather", map[string]string{
		"q": city,
	}))

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	return body
}

// GetDaily ...
func (ow *Provider) GetDaily(city string, maxDays uint) []byte {
	if maxDays < 1 || maxDays > 16 {
		panic("Max days must be between 1 and 16.")
	}

	res, err := http.Get(ow.GeneratePath("forecast/daily", map[string]string{
		"q":   city,
		"cnt": strconv.Itoa(int(maxDays)),
	}))

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	return body
}

// GeneratePath ...
func (ow *Provider) GeneratePath(endpoint string, urlParams map[string]string) string {
	params := url.Values{}
	params.Add("units", "metric")
	params.Add("APPID", ow.AppKey)

	for k, v := range urlParams {
		params.Add(k, v)
	}

	p := strings.TrimRight(ow.Path, "/") + "/" + strings.TrimLeft(endpoint, "/")

	return p + "?" + params.Encode()
}
