package db

import (
	"gopkg.in/mgo.v2"
)

// Insert generic function
func Insert(model interface{}, sess *mgo.Session, dbName interface{}) error {
	var ok bool
	var name string
	if sess == nil {
		sess = Database.Session
	}
	if dbName == nil {
		name = Name
	}
	if name, ok = dbName.(string); !ok {
		return UnknownDatabaseNameTypeError{}
	}

	cloneSess := sess.Clone()
	defer cloneSess.Close()

	switch model.(type) {
	case Product:
		e := cloneSess.DB(name).C("products").Insert(model)
		return e
	case Photo:
		e := cloneSess.DB(name).C("photos").Insert(model)
		return e
	case Shelflife:
		e := cloneSess.DB(name).C("shelflife").Insert(model)
		return e
	default:
		return UnknownModelError{}

	}
}

// Find generic query to db
func Find(query interface{}, t interface{}, sess *mgo.Session, dbName interface{}) (result interface{}, err error) {
	cloneSess := sess.Clone()
	defer cloneSess.Close()
	var (
		queryBSON []byte
	)

	var ok bool
	var name string
	if sess == nil {
		sess = Database.Session
	}
	if dbName == nil {
		name = Name
	}
	if name, ok = dbName.(string); !ok {
		return nil, UnknownDatabaseNameTypeError{}
	}

	switch t.(type) {
	case Photo:
		err = cloneSess.DB(name).C("photos").Find(queryBSON).One(&result)
		return
	case Product:
		err = cloneSess.DB(name).C("products").Find(queryBSON).One(&result)
		return
	case Shelflife:
		err = cloneSess.DB(name).C("shelflifes").Find(queryBSON).One(&result)
		return
	default:
		return nil, UnknownModelError{}
	}
}
