package storage

import (
	"errors"
	"gopkg.in/mgo.v2"
)

type storer struct {
	c *mgo.Collection
}

func (storer *storer) Store(data interface{}) error {
	return storer.c.Insert(data)
}

func NewStorage(
	addr string,
	database string,
	colletion string,
) (*storage, error) {
	session, err := mgo.Dial(addr)
	collection := session.DB(database).C(colletion)
	return &storer{
		c: colletion,
	}
}
