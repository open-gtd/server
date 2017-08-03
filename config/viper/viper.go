package viper

import (
	"time"

	"github.com/open-gtd/server/config"
	"github.com/spf13/viper"
)

type reader struct {
}

//New - creates new configuration reader builder
func New() ReaderBuilder {
	return &readerBuilder{}
}

//ReaderBuilder - configuration reader builder
type ReaderBuilder interface {
	FileName(fileName string) ReaderBuilder
	AddConfigPath(path string) ReaderBuilder
	Build() (config.Reader, error)
}

type readerBuilder struct {
	fileName string
	paths    []string
}

func (r *readerBuilder) FileName(fileName string) ReaderBuilder {
	r.fileName = fileName

	return r
}

func (r *readerBuilder) AddConfigPath(path string) ReaderBuilder {
	r.paths = append(r.paths, path)

	return r
}

func (r *readerBuilder) Build() (config.Reader, error) {
	viper.SetConfigName(r.fileName)

	for _, path := range r.paths {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return reader{}, nil
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
