package eventBus

type Topic string

type Bus interface {
	Subscribe(topic Topic, fn interface{}) error
	Publish(topic Topic, arg interface{})
}

type BusCollection interface {
	New(name Topic) Bus
	Get(name Topic) Bus
}
