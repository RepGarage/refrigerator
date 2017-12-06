package db

import (
	"os"
	"testing"

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
	var testProduct = Product{
		Name:      "Product",
		ProductID: 1,
	}
	err := Insert(testProduct, sess, dbName)
	if err != nil {
		t.Error(err)
	}
	c, e := sess.DB(dbName).C("products").Count()
	if e != nil {
		t.Error(e)
	}
	assert.Equal(t, 1, c, "products collection size should be 1")
	_err := sess.DB(dbName).C("products").DropCollection()
	if err != nil {
		t.Error(_err)
	}
}

func TestInsertPhoto(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testPhoto = Photo{
		Data:      "testphoto",
		ProductID: 1,
		Side:      100,
	}
	err := Insert(testPhoto, sess, dbName)
	if err != nil {
		t.Error(err)
	}
	c, e := sess.DB(dbName).C("photos").Count()
	if e != nil {
		t.Error(e)
	}
	assert.Equal(t, 1, c, "collection size should be 1")
	_err := sess.DB(dbName).C("photos").DropCollection()
	if err != nil {
		t.Error(_err)
	}
}

func TestInsertUnknown(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	type testUnknown struct {
		Name string
	}
	err := Insert(testUnknown{Name: "unknown"}, sess, dbName)
	assert.Equal(t, UnknownModelError{}, err, "should be returned UnknownModelError")
}

func TestFindProductByID(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testProduct = Product{
		Name:      "Product",
		ProductID: 1,
	}
	err := Insert(testProduct, sess, dbName)
	if err != nil {
		t.Error(err)
	}
	pInt, err := FindProductByID(testProduct.ProductID, sess, dbName)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Product", pInt.Name, "name should be Product")
	assert.Equal(t, 1, pInt.ProductID, "productID should be 1")

	pString, err := FindProductByID("1", sess, dbName)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Product", pString.Name, "name should be Product")
	assert.Equal(t, 1, pString.ProductID, "productID should be 1")

	_err := sess.DB(dbName).C("products").DropCollection()
	if err != nil {
		t.Error(_err)
	}
}

func TestFindPhotoByProductIDAndSide(t *testing.T) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		t.Error(e)
	}
	var testPhoto = Photo{
		Data:      "testphoto",
		ProductID: 1,
		Side:      100,
	}
	err := sess.DB(dbName).C("photos").Insert(testPhoto)
	if err != nil {
		t.Error(err)
	}

	// ProductID and Side int
	p, err := FindPhotoByProductID(testPhoto.ProductID, testPhoto.Side, sess, dbName)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "testphoto", p.Data, "data should be testphoto")
	assert.Equal(t, 1, p.ProductID, "products id should be 1")
	assert.Equal(t, 100, p.Side, "side should be 100")

	// ProdcutID string, Side int
	pSI, err := FindPhotoByProductID("1", testPhoto.Side, sess, dbName)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "testphoto", pSI.Data, "data should be testphoto")
	assert.Equal(t, 1, pSI.ProductID, "products id should be 1")
	assert.Equal(t, 100, pSI.Side, "side should be 100")

	// ProductID and Side string
	pSS, err := FindPhotoByProductID("1", "100", sess, dbName)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "testphoto", pSS.Data, "data should be testphoto")
	assert.Equal(t, 1, pSS.ProductID, "products id should be 1")
	assert.Equal(t, 100, pSS.Side, "side should be 100")

	_err := sess.DB(dbName).C("photos").DropCollection()
	if err != nil {
		t.Error(_err)
	}
}

func BenchmarkInsert(b *testing.B) {
	sess, e := mgo.Dial(mongoURL)
	if e != nil {
		b.Error(e)
	}
	for i := 0; i < b.N; i++ {
		var testProduct = Product{
			Name:      "Product",
			ProductID: i,
		}
		err := Insert(testProduct, sess, dbName)
		if err != nil {
			b.Error(err)
		}
	}
}
