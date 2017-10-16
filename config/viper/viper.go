package viper

import (
	"time"

	"github.com/open-gtd/server/config"
	"github.com/spf13/viper"
)

type Configuration interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
}

var v Configuration = &viper.Viper{}

type reader struct {
}

func New() config.Reader {
	return reader{}
}

func (r reader) Get(key string) interface{} {
	return v.Get(key)
}

func (r reader) GetBool(key string) bool {
	return v.GetBool(key)
}
func (r reader) GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}
func (r reader) GetInt(key string) int {
	return v.GetInt(key)
}
func (r reader) GetString(key string) string {
	return v.GetString(key)
}
func (r reader) GetStringMap(key string) map[string]interface{} {
	return v.GetStringMap(key)
}
func (r reader) GetStringMapString(key string) map[string]string {
	return v.GetStringMapString(key)
}
func (r reader) GetStringSlice(key string) []string {
	return v.GetStringSlice(key)
}
func (r reader) GetTime(key string) time.Time {
	return v.GetTime(key)
}
func (r reader) GetDuration(key string) time.Duration {
	return v.GetDuration(key)
}
func (r reader) IsSet(key string) bool {
	return v.IsSet(key)
}
