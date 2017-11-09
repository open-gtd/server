package config

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert.Nil(t, NullReader{}.Get(""))
}

func TestGetBool(t *testing.T) {
	assert.Equal(t, false, NullReader{}.GetBool(""))
}

func TestGetFloat64(t *testing.T) {
	assert.Equal(t, float64(0), NullReader{}.GetFloat64(""))
}

func TestGetInt(t *testing.T) {
	assert.Equal(t, 0, NullReader{}.GetInt(""))
}

func TestGetString(t *testing.T) {
	assert.Equal(t, "", NullReader{}.GetString(""))
}

func TestGetStringMap(t *testing.T) {
	assert.Equal(t, make(map[string]interface{}), NullReader{}.GetStringMap(""))
}

func TestGetStringMapString(t *testing.T) {
	assert.Equal(t, make(map[string]string), NullReader{}.GetStringMapString(""))
}

func TestGetStringSlice(t *testing.T) {
	assert.Equal(t, []string{}, NullReader{}.GetStringSlice(""))
}

func TestGetTime(t *testing.T) {
	assert.Equal(t, time.Time{}, NullReader{}.GetTime(""))
}

func TestGetDuration(t *testing.T) {
	assert.Equal(t, time.Duration(0), NullReader{}.GetDuration(""))
}

func TestIsSet(t *testing.T) {
	assert.False(t, NullReader{}.IsSet(""))
}
