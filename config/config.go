package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// OpenWeatherConfig ...
type OpenWeatherConfig struct {
	Path   string
	AppKey string
}

// LoggerConfig ...
type LoggerConfig struct {
	File string
}

// ResponseConfig ...
type ResponseConfig struct {
	Timeout time.Duration
}

// Config ...
type Config struct {
	OpenWeather OpenWeatherConfig
	Response    ResponseConfig
}

var configCache *Config

// Get ...
func Get() *Config {
	if configCache == nil {
		panic("App config is not loaded.")
	}

	return configCache
}

// Load ...
func Load() {
	if configCache != nil {
		panic("You can only load config once.")
	}

	configCache = &Config{
		OpenWeather: OpenWeatherConfig{
			Path:   getEnvString("OPENWEATHER_PATH", ""),
			AppKey: getEnvString("OPENWEATHER_APP_KEY", ""),
		},
		Response: ResponseConfig{
			Timeout: time.Millisecond * time.Duration(getEnvInt("RESPONSE_TIMEOUT", 1000)),
		},
	}
}

func getEnvString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if valueInt, err := strconv.Atoi(value); err == nil {
			return valueInt
		}

		panic(fmt.Sprintf("Cannot parse env %q value to int.", value))
	}

	return defaultVal
}
