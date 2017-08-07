package eventBus

var b Bus = NullBus{}

func SetBus(bus Bus) {
	b = bus
}

func GetBus() Bus {
	return b
}
