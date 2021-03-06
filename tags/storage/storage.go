package storage

const (
	NotFoundError  = "not found"
	NotUniqueError = "not unique"
)

type Dao interface {
	Insert(tag Tag) error
	Update(name string, tag Tag) error
	Delete(name string) error
	GetList() ([]Tag, error)
	Get(name string) (Tag, error)
}
