package api

type (
	TypeDescriptor string
	Name           string
)

const (
	Label   TypeDescriptor = "label"
	Area    TypeDescriptor = "area"
	Contact TypeDescriptor = "contact"
)

type Tag struct {
	Name Name           `json: "name"`
	Type TypeDescriptor `json: "type"`
}
