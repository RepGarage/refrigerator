package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/centrypoint/refrigerator/back/go/main/db"
)

var mongoURL = os.Getenv("MONGO_URL")

// DatabaseInterface database interface
type DatabaseInterface interface {
	Connect(string) error
	FetchProductFromDBByID(interface{}) (db.Product, error)
	FetchPhotoFromDBByProductID(productID interface{}, side interface{}) (db.Photo, error)
	InsertProductToDB(prod db.Product) error
	InsertPhotoToDB(photo db.Photo) error
}

// Database DatabaseInterface implementation
type Database struct {
	Session  *mgo.Session
	Database *mgo.Database
}

type mongoInterface interface {
	Dial(string) (*mgo.Session, error)
}

// Connect create connection
func (d *Database) Connect(dbName string) error {
	session, error := mgo.Dial(mongoURL)
	if error != nil {
		return error
	}
	d.Session = session
	database := session.DB(dbName)
	d.Database = database
	return nil
}

// FetchProductFromDBByID func
func (d *Database) FetchProductFromDBByID(id interface{}) (db.Product, error) {
	return db.FindProductByID(id, d.Session, d.Database.Name)
}

// FetchPhotoFromDBByProductID func
func (d *Database) FetchPhotoFromDBByProductID(productID interface{}, side interface{}) (db.Photo, error) {
	return db.FindPhotoByProductID(productID, side, d.Session, d.Database.Name)
}

// InsertProductToDB func
func (d *Database) InsertProductToDB(prod db.Product) error {
	log.Printf("Inserting product %+v to db\n", prod)
	return db.Insert(prod, d.Session, d.Database.Name)
}

// InsertPhotoToDB func
func (d *Database) InsertPhotoToDB(photo db.Photo) error {
	log.Printf("Inserting photo %+v to db\n", photo)
	return db.Insert(photo, d.Session, d.Database.Name)
}
