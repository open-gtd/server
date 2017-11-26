package storage

const (
	Empty   uint8 = iota
	Label
	Area
	Contact
)

type Tag struct {
	Name string
	Type uint8
}

var EmptyTag = Tag{}
