package factories

import (
	"errors"

	commonsPool "github.com/jolestar/go-commons-pool"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/tags/storage"
	"github.com/open-gtd/server/tags/storage/mongo"
	"github.com/open-gtd/server/pool"
)

var c = config.GetReader()

var createDao = mongo.CreateDao

var daoPool pool.ObjectPool = commonsPool.NewObjectPool(
	commonsPool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {

			host := c.GetString("tags.database.host")
			database := c.GetString("tags.database.database")
			collection := c.GetString("tags.database.collection")
			dao, err := createDao(
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
	commonsPool.NewDefaultPoolConfig(),
)

func Dao() (storage.Dao, error) {
	obj, err := daoPool.BorrowObject()
	if err != nil {
		return nil, err
	}

	dao, ok := obj.(storage.Dao)
	if ok != true {
		return nil, errors.New("object in commonsPool is in wrong type, should never happen")
	}

	return dao, nil
}
