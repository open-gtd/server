package factories

import (
	"errors"

	"github.com/jolestar/go-commons-pool"
	"github.com/open-gtd/server/tags/config"
	"github.com/open-gtd/server/tags/storage"
	"github.com/open-gtd/server/tags/storage/mongo"
)

var p = pool.NewObjectPool(
	pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {

			c := config.Get()

			dao, err := mongo.CreateDao(
				c.GetString("tags.database.host"),
				c.GetString("tags.database.database"),
				c.GetString("tags.database.collection"),
			)
			if err != nil {
				return nil, err
			}
			return dao, nil
		},
	),
	pool.NewDefaultPoolConfig(),
)

func Dao() (storage.Dao, error) {
	obj, err := p.BorrowObject()
	if err != nil {
		return nil, err
	}

	dao, ok := obj.(storage.Dao)
	if ok != true {
		return nil, errors.New("object in pool is in wrong type, should never happen")
	}

	return dao, nil
}

func returnDao(dao storage.Dao) error {
	return p.ReturnObject(dao)
}
