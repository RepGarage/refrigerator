package main

import (
	"./db"
	"gopkg.in/mgo.v2"
)

// Db database struct
type Db struct {
	session  *mgo.Session
	database *mgo.Database
}

// Connect create connection
func (d *Db) Connect() error {
	session, error := mgo.Dial("localhost")
	if error != nil {
		return error
	}
	d.session = session
	database := session.DB("refrigerator-dev")
	d.database = database
	return nil
}

// FetchProductFromDBByID func
func (d *Db) FetchProductFromDBByID(id int) (db.Product, error) {
	p, e := db.FindProductByID(id, d.session, d.database.Name)
	return p, e
}

// FetchPhotoFromDBByProductID func
func (d *Db) FetchPhotoFromDBByProductID(productID int) (db.Photo, error) {
	return db.FindPhotoByProductID(productID, d.session, d.database.Name)
}

// InsertProductToDB func
func (d *Db) InsertProductToDB(prod db.Product) error {
	return db.Insert(prod, d.session, d.database.Name)
}

// InsertPhotoToDB func
func (d *Db) InsertPhotoToDB(photo db.Photo) error {
	return db.Insert(photo, d.session, d.database.Name)
}
