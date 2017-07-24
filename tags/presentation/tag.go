package presentation

type (
	TypeDescriptor string
	Name           string
)

const (
	EmptyType TypeDescriptor = ""
	Label     TypeDescriptor = "label"
	Area      TypeDescriptor = "area"
	Contact   TypeDescriptor = "contact"
)

type Tag struct {
	Name Name           `json: "name"`
	Type TypeDescriptor `json: "type"`
}
