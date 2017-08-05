package tags

const (
	EmptyType string = ""
	Label     string = "label"
	Area      string = "area"
	Contact   string = "contact"
)

var AllowedTypes = []string{Label, Area, Contact}

type TagId struct {
	Name string `json:"name"`
}

type Tag struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreatedTag Tag

type ModifiedTag struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	OriginalName string `json:"original_name"`
}

type DeletedTag TagId
