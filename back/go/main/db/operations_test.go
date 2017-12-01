package db_test

import (
	"testing"

	"bitbucket.org/ngoncharov/refrigerator/db"

	"gopkg.in/mgo.v2"
)

const dbName = "refrigerator-test"

func TestInsert(t *testing.T) {
	sess, e := mgo.Dial("localhost")
	if e != nil {
		t.Error(e)
	}
	var testProduct = db.Product{
		Name:      "Product",
		ProductID: 1,
	}
	err := db.Insert(testProduct, sess, dbName)
	if err != nil {
		t.Error(err)
	}
	_err := sess.DB(dbName).C("products").DropCollection()
	if err != nil {
		t.Error(_err)
	}
}

func TestFindProductByID(t *testing.T) {
	sess, e := mgo.Dial("localhost")
	if e != nil {
		t.Error(e)
	}
	var testProduct = db.Product{
		Name:      "Product",
		ProductID: 1,
	}
	err := db.Insert(testProduct, sess, dbName)
	if err != nil {
		t.Error(err)
	}
	p, err := db.FindProductByID(testProduct.ProductID, sess, dbName)
	if err != nil {
		t.Error(err)
	}
	if len(p.Name) == 0 {
		t.Errorf("product name nil\n")
	}

	_err := sess.DB(dbName).C("products").DropCollection()
	if err != nil {
		t.Error(_err)
	}
}

func BenchmarkInsert(b *testing.B) {
	sess, e := mgo.Dial("localhost")
	if e != nil {
		b.Error(e)
	}
	for i := 0; i < b.N; i++ {
		var testProduct = db.Product{
			Name:      "Product",
			ProductID: i,
		}
		err := db.Insert(testProduct, sess, dbName)
		if err != nil {
			b.Error(err)
		}
	}
}
