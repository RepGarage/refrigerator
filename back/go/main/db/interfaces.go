package db

import mgo "gopkg.in/mgo.v2"

// Inserter provide insert operation for db
type Inserter interface {
	Insert(model interface{}, sess *mgo.Session, dbName interface{}) error
}

// Finder provide db qieries
type Finder interface {
	Find(query interface{}, t interface{}, sess *mgo.Session, dbName interface{}) (result interface{}, err error)
}

// FindInserter implements both find and insert operations
type FindInserter interface {
	Insert(model interface{}, sess *mgo.Session, dbName interface{}) error
	Find(query interface{}, t interface{}, sess *mgo.Session, dbName interface{}) (result interface{}, err error)
}
