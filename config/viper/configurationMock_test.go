package viper

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type configurationMock struct {
	mock.Mock
}

func (c *configurationMock) Get(key string) interface{} {
	args := c.Called(key)
	return args.Get(0)
}

func (c *configurationMock) GetBool(key string) bool {
	args := c.Called(key)
	return args.Bool(0)
}

func (c *configurationMock) GetFloat64(key string) float64 {
	args := c.Called(key)
	return args.Get(0).(float64)
}

func (c *configurationMock) GetInt(key string) int {
	args := c.Called(key)
	return args.Int(0)
}

func (c *configurationMock) GetString(key string) string {
	args := c.Called(key)
	return args.String(0)
}

func (c *configurationMock) GetStringMap(key string) map[string]interface{} {
	args := c.Called(key)
	return args.Get(0).(map[string]interface{})
}

func (c *configurationMock) GetStringMapString(key string) map[string]string {
	args := c.Called(key)
	return args.Get(0).(map[string]string)
}

func (c *configurationMock) GetStringSlice(key string) []string {
	args := c.Called(key)
	return args.Get(0).([]string)
}

func (c *configurationMock) GetTime(key string) time.Time {
	args := c.Called(key)
	return args.Get(0).(time.Time)
}

func (c *configurationMock) GetDuration(key string) time.Duration {
	args := c.Called(key)
	return args.Get(0).(time.Duration)
}

func (c *configurationMock) IsSet(key string) bool {
	args := c.Called(key)
	return args.Bool(0)
}
