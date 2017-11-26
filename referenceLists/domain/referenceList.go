package domain

import (
	"github.com/open-gtd/server/tags/domain"
)

type (
	Name string
	Tags []domain.Tag
	Notes string
	Focus bool
)

type ReferenceListDescriptor struct {
	Name Name `json:"name"`
}

type ReferenceList struct {
	ReferenceListDescriptor
	Tags  Tags  `json:"tags"`
	Notes Notes `json:"notes"`
}
