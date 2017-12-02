package pool

type ObjectPool interface {
	BorrowObject() (interface{}, error)
	ReturnObject(object interface{}) error
}
