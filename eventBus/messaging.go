package eventBus

type Topic string
type Name string

var b Bus = nullBus{}

type BusHandler func(arg interface{})

type Bus interface {
	Subscribe(topic Topic, fn BusHandler) error
	Unsubscribe(topic Topic, handler BusHandler) error
	Publish(topic Topic, arg interface{})
}

func SetBus(bus Bus) {
	b = bus
}

func GetBus() Bus {
	return b
}

type nullBus struct{}

func (nullBus) Subscribe(topic Topic, fn BusHandler) error {
	return nil
}
func (nullBus) Unsubscribe(topic Topic, handler BusHandler) error {
	return nil
}
func (nullBus) Publish(topic Topic, arg interface{}) {}
