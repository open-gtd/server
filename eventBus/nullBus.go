package eventBus

type NullBus struct{}

func (NullBus) Subscribe(topic Topic, fn BusHandler) error {
	return nil
}
func (NullBus) Unsubscribe(topic Topic, handler BusHandler) error {
	return nil
}
func (NullBus) Publish(topic Topic, arg interface{}) {}
