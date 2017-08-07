package viper

import (
	"time"

	"github.com/open-gtd/server/config"
	"github.com/spf13/viper"
)

type reader struct {
}

func New() config.Reader {
	return reader{}
}

func (r reader) Get(key string) interface{} {
	return viper.Get(key)
}

func (r reader) GetBool(key string) bool {
	return viper.GetBool(key)
}
func (r reader) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}
func (r reader) GetInt(key string) int {
	return viper.GetInt(key)
}
func (r reader) GetString(key string) string {
	return viper.GetString(key)
}
func (r reader) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
func (r reader) GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
func (r reader) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
func (r reader) GetTime(key string) time.Time {
	return viper.GetTime(key)
}
func (r reader) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
func (r reader) IsSet(key string) bool {
	return viper.IsSet(key)
}
