package storage

import (
	"gopkg.in/mgo.v2"
)

type storage struct {
	c *mgo.Collection
}

func (s *storage) Store(data interface{}) error {
	return s.c.Insert(data)
}

func NewStorage(
	addr string,
	database string,
	colletion string,
) (*storage, error) {
	session, err := mgo.Dial(addr)
	if err != nil {
		return &storage{}, err
	}
	c := session.DB(database).C(colletion)
	return &storage{
		c: c,
	}, nil
}
