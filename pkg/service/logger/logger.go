package logger

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Inst ...
func Inst() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}

	return logger
}

// LoadConfig ...
func LoadConfig() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

// LogHTTPError ...
func LogHTTPError(r *http.Request, text string) {
	str, err := json.Marshal(struct {
		Path    string `json:"path"`
		Message string `json:"message"`
	}{
		r.URL.Path,
		text,
	})

	if err != nil {
		return
	}

	logger.Error(string(str))
}
