package config

import "github.com/open-gtd/server/config"

var reader config.Reader = config.NullReader{}

func RegisterReader(r config.Reader) {
	reader = r
}

func Get() config.Reader {
	return reader
}
