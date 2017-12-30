package db

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	url = os.Getenv("MONGO_URL")
	// Name database name
	Name = os.Getenv("DATABASE_NAME")
	// Database ref
	Database *mgo.Database
)

func init() {
	log.Fatal(Connect())
}

// Connect establish db connection
func Connect() (err error) {
	var sess *mgo.Session
	if sess, err = mgo.Dial(url); err != nil {
		return
	}

	Database = sess.DB(Name)
	return nil
}
