package config

import "time"

//Reader - represents ability to read configuration from file
type Reader interface {
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

type NullReader struct{}

func (r NullReader) Get(key string) interface{} {
	return nil
}
func (r NullReader) GetBool(key string) bool {
	return false
}
func (r NullReader) GetFloat64(key string) float64 {
	return 0
}
func (r NullReader) GetInt(key string) int {
	return 0
}
func (r NullReader) GetString(key string) string {
	return ""
}
func (r NullReader) GetStringMap(key string) map[string]interface{} {
	return make(map[string]interface{})
}
func (r NullReader) GetStringMapString(key string) map[string]string {
	return make(map[string]string)
}
func (r NullReader) GetStringSlice(key string) []string {
	return []string{}
}
func (r NullReader) GetTime(key string) time.Time {
	return time.Time{}
}
func (r NullReader) GetDuration(key string) time.Duration {
	return 0
}
func (r NullReader) IsSet(key string) bool {
	return false
}
