package reference

import tags "github.com/open-gtd/server/tags/api"

type (
	Name  string
	Tags  []tags.Tag
	Notes string
	Focus bool
)

type ReferenceListDescriptor struct {
	Name Name `json: "name"`
}

type ReferenceList struct {
	ReferenceListDescriptor
	Tags  Tags  `json: "tags"`
	Notes Notes `json: "notes"`
}
