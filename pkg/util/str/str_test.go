package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	isEmpty := IsEmpty("")

	assert.True(t, isEmpty)
}

func TestEmptyStringWithSpaces(t *testing.T) {
	isEmpty := IsEmpty("   ")

	assert.True(t, isEmpty)
}

func TestNotEmptyString(t *testing.T) {
	isEmpty := IsEmpty("привет")

	assert.False(t, isEmpty)
}
