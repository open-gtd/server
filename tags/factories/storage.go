package factories

import (
	"errors"

	"github.com/jolestar/go-commons-pool"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/tags/storage"
	"github.com/open-gtd/server/tags/storage/mongo"
)

var c = config.GetReader()

var p = pool.NewObjectPool(
	pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {

			host := c.GetString("tags.database.host")
			database := c.GetString("tags.database.database")
			collection := c.GetString("tags.database.collection")
			dao, err := mongo.CreateDao(
				host,
				database,
				collection,
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
