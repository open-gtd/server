package config

var r Reader = NullReader{}

func SetReader(reader Reader) {
	r = reader
}

func GetReader() Reader {
	return r
}
