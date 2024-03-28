package structs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStringSetting(t *testing.T) {
	name := "Name"
	value := "Value"

	setting := NewSetting(name, value)

	assert.Equal(t, name, setting.Name)
	assert.Equal(t, value, setting.Value)
}

func TestNewTimeSetting(t *testing.T) {
	name := "Name"
	value := time.Now()

	setting := NewSetting(name, value)

	assert.Equal(t, name, setting.Name)
	assert.Equal(t, value, setting.Value)
}

func TestNewBoolSetting(t *testing.T) {
	name := "Name"
	value := true

	setting := NewSetting(name, value)

	assert.Equal(t, name, setting.Name)
	assert.Equal(t, value, setting.Value)
}

func TestNewIntSetting(t *testing.T) {
	name := "Name"
	value := 42

	setting := NewSetting(name, value)

	assert.Equal(t, name, setting.Name)
	assert.Equal(t, value, setting.Value)
}
