package sse

var r Registerer = NullRegisterer{}

func SetRegisterer(registerer Registerer) {
	r = registerer
}

func GetRegistarer() Registerer {
	return r
}
