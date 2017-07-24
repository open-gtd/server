package mongo

import (
	dtags "github.com/open-gtd/server/domain/tags"
	"github.com/open-gtd/server/storage/tags"
	"labix.org/v2/mgo"
)

type storage struct {
	session    *mgo.Session
	database   string
	collection string
}

const getAllQuery = ""

func (s storage) GetAll() ([]dtags.Tag, error) {
	var result []dtags.Tag
	c := s.session.DB(s.database).C(s.collection)

	err := c.Find(getAllQuery).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateStorage(url string) (tags.Storage, error) {
	s := storage{}
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	s.session = session

	return s, nil
}
