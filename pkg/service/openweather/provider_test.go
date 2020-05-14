package openweather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider_GeneratePath(t *testing.T) {
	provider := Provider{
		Path:   "https://lorem.ipsum",
		AppKey: "MY-APP-KEY",
	}

	actual := provider.GeneratePath("users-list", map[string]string{
		"limit": "10",
		"page":  "1",
	})

	expected := "https://lorem.ipsum/users-list?APPID=MY-APP-KEY&limit=10&page=1&units=metric"

	assert.Equal(t, expected, actual)
}

func TestProvider_GeneratePath_WithExtraSlashes(t *testing.T) {
	provider := Provider{
		Path:   "https://lorem.ipsum/",
		AppKey: "MY-APP-KEY",
	}

	actual := provider.GeneratePath("/users-list", map[string]string{
		"limit": "10",
		"page":  "1",
	})

	expected := "https://lorem.ipsum/users-list?APPID=MY-APP-KEY&limit=10&page=1&units=metric"

	assert.Equal(t, expected, actual)
}
