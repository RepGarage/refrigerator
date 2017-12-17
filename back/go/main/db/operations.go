package db

import (
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Prepare ensure indexes
// func Prepare(sess *mgo.Session, dbName) error {
// 	clone := sess.Clone()
// 	defer clone.Close()
// 	clone.DB(dbName).C("products").EnsureIndexKey("_id")
// }

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
func FindPhotoByProductID(id interface{}, side interface{}, sess *mgo.Session, dbName string) (Photo, error) {
	cloneSess := sess.Clone()
	defer cloneSess.Close()
	var result Photo
	var iID int
	var iSide int
	// Converts
	switch v := id.(type) {
	case string:
		iID, _ = strconv.Atoi(v)
		break
	case int:
		iID = v
		break
	}

	switch v := side.(type) {
	case string:
		iSide, _ = strconv.Atoi(v)
		break
	case int:
		iSide = v
		break
	}

	e := cloneSess.DB(dbName).C("photos").Find(bson.M{"_id": iID, "side": iSide}).One(&result)
	return result, e
}

// FindProductByID func
func FindProductByID(id interface{}, sess *mgo.Session, dbName string) (Product, error) {
	cloneSess := sess.Clone()
	defer cloneSess.Close()
	var result Product
	var iID int
	// Converts
	switch v := id.(type) {
	case string:
		iID, _ = strconv.Atoi(v)
		break
	case int:
		iID = v
		break
	}
	e := cloneSess.DB(dbName).C("products").Find(bson.M{"_id": iID}).One(&result)
	return result, e
}
