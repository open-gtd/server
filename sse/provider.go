package sse

var registerer Registerer = NullRegisterer{}

func SetRegisterer(r Registerer) {
	registerer = r
}

func GetRegisterer() Registerer {
	return registerer
}
