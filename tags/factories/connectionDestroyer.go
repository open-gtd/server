package factories

import "github.com/open-gtd/server/tags/storage"

type connectionDestroyer struct {
	conn storage.Dao
}

func (c connectionDestroyer) Destroy() error {
	return daoPool.ReturnObject(c.conn)
}
