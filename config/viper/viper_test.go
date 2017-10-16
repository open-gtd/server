package viper

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"time"
)

func TestReader_Get(t *testing.T) {

	const value = "c"
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("Get", key).Return(value)

	sut := New()

	result := sut.Get(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetBool(t *testing.T) {

	const value = true
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetBool", key).Return(value)

	sut := New()

	result := sut.GetBool(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetFloat64(t *testing.T) {

	const value = 3.45
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetFloat64", key).Return(value)

	sut := New()

	result := sut.GetFloat64(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetInt(t *testing.T) {

	const value = 5
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetInt", key).Return(value)

	sut := New()

	result := sut.GetInt(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetString(t *testing.T) {

	const value = "str"
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetString", key).Return(value)

	sut := New()

	result := sut.GetString(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetStringMap(t *testing.T) {

	var value = map[string]interface{}{}
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetStringMap", key).Return(value)

	sut := New()

	result := sut.GetStringMap(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetStringMapString(t *testing.T) {

	var value = map[string]string{}
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetStringMapString", key).Return(value)

	sut := New()

	result := sut.GetStringMapString(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetStringSlice(t *testing.T) {

	var value = []string{}
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetStringSlice", key).Return(value)

	sut := New()

	result := sut.GetStringSlice(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetTime(t *testing.T) {

	var value = time.Time{}
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetTime", key).Return(value)

	sut := New()

	result := sut.GetTime(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_GetDuration(t *testing.T) {

	var value = time.Duration(15)
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("GetDuration", key).Return(value)

	sut := New()

	result := sut.GetDuration(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}

func TestReader_IsSet(t *testing.T) {

	const value = true
	const key = "k"

	configurationMock := &configurationMock{}
	v = configurationMock

	configurationMock.On("IsSet", key).Return(value)

	sut := New()

	result := sut.IsSet(key)

	configurationMock.AssertExpectations(t)
	assert.Equal(t, result, value)
}
