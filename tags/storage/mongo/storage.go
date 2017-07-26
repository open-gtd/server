package mongo

import (
	"errors"

	"github.com/open-gtd/server/tags/storage"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const notFoundError = "not found"

type dao struct {
	session    *mgo.Session
	database   string
	collection string
}

func (s dao) GetList() ([]storage.Tag, error) {
	var result []storage.Tag
	c := s.session.DB(s.database).C(s.collection)

	err := c.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s dao) Get(name string) (storage.Tag, error) {
	c := s.session.DB(s.database).C(s.collection)

	tag := storage.Tag{}
	err := c.Find(bson.M{"name": name}).One(&tag)
	if err != nil {
		if err.Error() == notFoundError {
			err = errors.New(storage.NotFoundError)
		}

		return storage.EmptyTag, err
	}

	return tag, nil
}

func (s dao) Update(name string, tag storage.Tag) error {
	c := s.session.DB(s.database).C(s.collection)

	err := c.Update(bson.M{"name": name}, tag)
	if err != nil {
		if err.Error() == notFoundError {
			err = errors.New(storage.NotFoundError)
		}

		return err
	}

	return nil
}

func (s dao) Insert(tag storage.Tag) error {
	c := s.session.DB(s.database).C(s.collection)

	err := c.Insert(tag)
	if err != nil {
		return err
	}

	return nil
}

func (s dao) Delete(name string) error {
	c := s.session.DB(s.database).C(s.collection)

	err := c.Remove(bson.M{"name": name})
	if err != nil {
		if err.Error() == notFoundError {
			return errors.New(storage.NotFoundError)
		}

		return err
	}

	return nil
}

func CreateDao(url string) (storage.Dao, error) {
	s := dao{}
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	s.session = session
	s.database = "open-gtd"
	s.collection = "tags"

	return s, nil
}
