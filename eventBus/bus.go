package eventBus

type Topic string
type Name string

type BusHandler func(arg interface{})

type Bus interface {
	Subscribe(topic Topic, fn BusHandler) error
	Unsubscribe(topic Topic, handler BusHandler) error
	Publish(topic Topic, arg interface{})
}

type NullBus struct{}

func (NullBus) Subscribe(topic Topic, fn BusHandler) error {
	return nil
}
func (NullBus) Unsubscribe(topic Topic, handler BusHandler) error {
	return nil
}
func (NullBus) Publish(topic Topic, arg interface{}) {}
