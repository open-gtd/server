package domain

import "github.com/open-gtd/server/tags/presentation"

type (
	Name  string
	Tags  []presentation.Tag
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
