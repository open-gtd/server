package api

var r Registerer = NullRegisterer{}

func SetRegisterer(registerer Registerer) {
	r = registerer
}

func GetRegisterer() Registerer {
	return r
}
