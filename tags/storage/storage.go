package storage

import "github.com/open-gtd/server/domain/tags"

type Storage interface {
	GetAll() ([]tags.Tag, error)
}
