package db_test

import (
	"os"
	"testing"

	"github.com/centrypoint/refrigerator/back/go/main/db"
	"github.com/stretchr/testify/assert"

	"gopkg.in/mgo.v2"
)

const dbName = "refrigerator-test"

var mongoURL = os.Getenv("MONGO_URL")

func TestInsertProduct(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testProduct = db.Product{
		Name:      "Product",
		ProductID: 1,
	}
	err := db.Insert(testProduct, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("products").DropCollection()
		t.Error(err)
	}
	c, e := sess.DB(dbName).C("products").Count()
	if e != nil {
		sess.DB(dbName).C("products").DropCollection()
		t.Error(e)
	}
	assert.Equal(t, 1, c, "products collection size should be 1")
	sess.DB(dbName).C("products").DropCollection()
}

func TestInsertShelflife(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testShelflife = db.Shelflife{
		ProductID: 1,
		Data:      180,
	}
	err := db.Insert(testShelflife, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("shelflife").DropCollection()
		t.Error(err)
	}
	c, e := sess.DB(dbName).C("shelflife").Count()
	if e != nil {
		sess.DB(dbName).C("shelflife").DropCollection()
		t.Error(e)
	}
	assert.Equal(t, 1, c, "shelflife collection size should be 1")
	sess.DB(dbName).C("shelflife").DropCollection()
}

func TestInsertPhoto(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testPhoto = db.Photo{
		Data:      "testphoto",
		ProductID: 1,
		Side:      100,
	}
	err := db.Insert(testPhoto, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("photos").DropCollection()
		t.Error(err)
	}
	c, e := sess.DB(dbName).C("photos").Count()
	if e != nil {
		sess.DB(dbName).C("photos").DropCollection()
		t.Error(e)
	}
	assert.Equal(t, 1, c, "collection size should be 1")
	sess.DB(dbName).C("photos").DropCollection()
}

func TestInsertUnknown(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	type testUnknown struct {
		Name string
	}
	err := db.Insert(testUnknown{Name: "unknown"}, sess, dbName)
	assert.Equal(t, db.UnknownModelError{}, err, "should be returned UnknownModelError")
}

func TestFindProductByID(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testProduct = db.Product{
		Name:      "Product",
		ProductID: 1,
	}
	err := db.Insert(testProduct, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("products").DropCollection()
		t.Error(err)
	}
	pInt, err := db.FindProductByID(testProduct.ProductID, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("products").DropCollection()
		t.Error(err)
	}

	assert.Equal(t, "Product", pInt.Name, "name should be Product")
	assert.Equal(t, 1, pInt.ProductID, "productID should be 1")

	pString, err := db.FindProductByID("1", sess, dbName)
	if err != nil {
		sess.DB(dbName).C("products").DropCollection()
		t.Error(err)
	}

	assert.Equal(t, "Product", pString.Name, "name should be Product")
	assert.Equal(t, 1, pString.ProductID, "productID should be 1")

	sess.DB(dbName).C("products").DropCollection()
}

func TestFindShelflifeByProductID(t *testing.T) {
	var session *mgo.Session
	var err error
	var shelf db.Shelflife
	if session, err = mgo.Dial(mongoURL); err != nil {
		t.Fatal(err)
	}
	var testShelflife = db.Shelflife{
		ProductID: 1,
		Data:      180,
	}
	if err = db.Insert(testShelflife, session, dbName); err != nil {
		t.Fatal(err)
	}
	if shelf, err = db.FindShelflifeByProductID(testShelflife.ProductID, session, dbName); err != nil {
		session.DB(dbName).C("shelflife").DropCollection()
		t.Fatal(err)
	}

	assert.Equal(t, 180, shelf.Data)
	assert.Equal(t, 1, shelf.ProductID, "productID should be 1")

	if shelf, err = db.FindShelflifeByProductID("1", session, dbName); err != nil {
		session.DB(dbName).C("shelflife").DropCollection()
		t.Fatal(err)
	}

	assert.Equal(t, 180, shelf.Data)
	assert.Equal(t, 1, shelf.ProductID, "productID should be 1")

	session.DB(dbName).C("shelflife").DropCollection()
}

func TestFindPhotoByProductIDAndSide(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testPhoto = db.Photo{
		Data:      "testphoto",
		ProductID: 1,
		Side:      100,
	}
	err := sess.DB(dbName).C("photos").Insert(testPhoto)
	if err != nil {
		sess.DB(dbName).C("photos").DropCollection()
		t.Error(err)
	}

	// ProductID and Side int
	p, err := db.FindPhotoByProductID(testPhoto.ProductID, testPhoto.Side, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("photos").DropCollection()
		t.Error(err)
	}

	assert.Equal(t, "testphoto", p.Data, "data should be testphoto")
	assert.Equal(t, 1, p.ProductID, "products id should be 1")
	assert.Equal(t, 100, p.Side, "side should be 100")

	// ProdcutID string, Side int
	pSI, err := db.FindPhotoByProductID("1", testPhoto.Side, sess, dbName)
	if err != nil {
		sess.DB(dbName).C("photos").DropCollection()
		t.Error(err)
	}

	assert.Equal(t, "testphoto", pSI.Data, "data should be testphoto")
	assert.Equal(t, 1, pSI.ProductID, "products id should be 1")
	assert.Equal(t, 100, pSI.Side, "side should be 100")

	// ProductID and Side string
	pSS, err := db.FindPhotoByProductID("1", "100", sess, dbName)
	if err != nil {
		sess.DB(dbName).C("photos").DropCollection()
		t.Error(err)
	}

	assert.Equal(t, "testphoto", pSS.Data, "data should be testphoto")
	assert.Equal(t, 1, pSS.ProductID, "products id should be 1")
	assert.Equal(t, 100, pSS.Side, "side should be 100")

	sess.DB(dbName).C("photos").DropCollection()
}

func BenchmarkInsert(b *testing.B) {
	sess, e := mgo.Dial(mongoURL)
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
