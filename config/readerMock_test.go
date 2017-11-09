package config

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type readerMock struct {
	mock.Mock
}

func (r *readerMock) Get(key string) interface{} {
	args := r.Called(key)
	return  args.Get(0)
}

func (r *readerMock) GetBool(key string) bool {
	args := r.Called(key)
	return  args.Bool(0)
}

func (r *readerMock) GetFloat64(key string) float64 {
	args := r.Called(key)
	return  args.Get(0).(float64)
}

func (r *readerMock) GetInt(key string) int {
	args := r.Called(key)
	return  args.Int(0)
}

func (r *readerMock) GetString(key string) string {
	args := r.Called(key)
	return  args.String(0)
}

func (r *readerMock) GetStringMap(key string) map[string]interface{} {
	args := r.Called(key)
	return  args.Get(0).(map[string]interface{})
}

func (r *readerMock) GetStringMapString(key string) map[string]string {
	args := r.Called(key)
	return  args.Get(0).(map[string]string)
}

func (r *readerMock) GetStringSlice(key string) []string {
	args := r.Called(key)
	return  args.Get(0).([]string)
}

func (r *readerMock) GetTime(key string) time.Time {
	args := r.Called(key)
	return  args.Get(0).(time.Time)
}

func (r *readerMock) GetDuration(key string) time.Duration {
	args := r.Called(key)
	return  args.Get(0).(time.Duration)
}

func (r *readerMock) IsSet(key string) bool {
	args := r.Called(key)
	return  args.Bool(0)
}
