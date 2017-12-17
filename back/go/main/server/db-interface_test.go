package server_test

import (
	"testing"

	"github.com/centrypoint/refrigerator/back/go/main/db"
	server "github.com/centrypoint/refrigerator/back/go/main/server"
	"github.com/stretchr/testify/assert"
)

var (
	dbi    server.Database
	dbName = "ref-go-test"
)

func TestConnect(t *testing.T) {
	err := dbi.Connect(dbName)
	if err != nil {
		t.Error(err)
	}
}

func TestFetchProductFromDBbyID(t *testing.T) {
	err := dbi.Connect(dbName)
	if err != nil {
		t.Fatal(err)
	}
	_, e := dbi.FetchProductFromDBByID(1)
	assert.Equal(t, "not found", e.Error(), "should not find product")
	dbi.Database.C("products").Insert(db.Product{Name: "product", ProductID: 1})

	p, e := dbi.FetchProductFromDBByID(1)
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, 1, p.ProductID, "productID should be 1")
	assert.Equal(t, "product", p.Name, "product name should be product")

	_e := dbi.Database.C("products").DropCollection()
	if _e != nil {
		t.Fatal(_e)
	}
}

func TestFetchPhotoFromDBByProductID(t *testing.T) {
	var err error
	var photo db.Photo
	err = dbi.Connect(dbName)
	if err != nil {
		t.Fatal(err)
	}
	_, err = dbi.FetchPhotoFromDBByProductID(1, 100)
	if err == nil {
		t.Fatalf("error is nil")
	}
	assert.Equal(t, "not found", err.Error(), "error should be not found")

	err = dbi.Database.C("photos").Insert(db.Photo{Data: "photo", ProductID: 1, Side: 100})
	if err != nil {
		t.Fatal(err)
	}

	photo, err = dbi.FetchPhotoFromDBByProductID(1, 100)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, photo.ProductID)
	assert.Equal(t, 100, photo.Side)
	assert.Equal(t, "photo", photo.Data)

	err = dbi.Database.C("photos").DropCollection()
	if err != nil {
		t.Fatal(err)
	}
}
