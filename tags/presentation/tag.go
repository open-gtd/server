package presentation

const (
	EmptyType string = ""
	Label     string = "label"
	Area      string = "area"
	Contact   string = "contact"
)

var AllowedTypes = []string{Label, Area, Contact}

type Tag struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
