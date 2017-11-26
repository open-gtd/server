package eventBus

type Topic string
type Name string

type BusHandler func(arg interface{})

type Bus interface {
	Subscribe(topic Topic, fn BusHandler) error
	Unsubscribe(topic Topic, handler BusHandler) error
	Publish(topic Topic, arg interface{})
}
