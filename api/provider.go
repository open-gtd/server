package api

var r Registerer = NullRegisterer{}

func SetRegisterer(reader Registerer) {
	r = reader
}

func GetRegisterer() Registerer {
	return r
}
