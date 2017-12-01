package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert generic function
func Insert(model interface{}, sess *mgo.Session, dbName string) error {
	cloneSess := sess.Clone()
	defer cloneSess.Close()
	switch model.(type) {
	case Product:
		e := cloneSess.DB(dbName).C("products").Insert(model)
		return e
	case Photo:
		e := cloneSess.DB(dbName).C("photos").Insert(model)
		return e
	default:
		return UnknownModelError{}
	}
}

// FindPhotoByProductID func
func FindPhotoByProductID(id int, sess *mgo.Session, dbName string) (Photo, error) {
	cloneSess := sess.Clone()
	defer cloneSess.Close()
	var result Photo
	e := cloneSess.DB(dbName).C("photos").Find(bson.M{"product_id": id}).One(&result)
	return result, e
}

// FindProductByID func
func FindProductByID(id int, sess *mgo.Session, dbName string) (Product, error) {
	cloneSess := sess.Clone()
	defer cloneSess.Close()
	var result Product
	e := cloneSess.DB(dbName).C("products").Find(bson.M{"product_id": id}).One(&result)
	return result, e
}
