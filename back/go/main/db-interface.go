package main

import (
	"os"

	"gopkg.in/mgo.v2"

	"github.com/RepGarage/refrigerator/back/go/main/db"
)

var mongoURL = os.Getenv("MONGO_URL")

// DatabaseInterface database struct
type DatabaseInterface struct {
	session  *mgo.Session
	database *mgo.Database
}

// Connect create connection
func (d *DatabaseInterface) Connect() error {
	session, error := mgo.Dial(mongoURL)
	if error != nil {
		return error
	}
	d.session = session
	database := session.DB("refrigerator-dev")
	d.database = database
	return nil
}

// FetchProductFromDBByID func
func (d *DatabaseInterface) FetchProductFromDBByID(id interface{}) (db.Product, error) {
	p, e := db.FindProductByID(id, d.session, d.database.Name)
	return p, e
}

// FetchPhotoFromDBByProductID func
func (d *DatabaseInterface) FetchPhotoFromDBByProductID(productID interface{}, side interface{}) (db.Photo, error) {
	return db.FindPhotoByProductID(productID, side, d.session, d.database.Name)
}

// InsertProductToDB func
func (d *DatabaseInterface) InsertProductToDB(prod db.Product) error {
	return db.Insert(prod, d.session, d.database.Name)
}

// InsertPhotoToDB func
func (d *DatabaseInterface) InsertPhotoToDB(photo db.Photo) error {
	return db.Insert(photo, d.session, d.database.Name)
}
