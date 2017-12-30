package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/centrypoint/refrigerator/back/go/main/db"
)

// ProductsAPI type
type ProductsAPI struct{}

// GetProductsHandler func
func (p *ProductsAPI) GetProductsHandler(
	api GetProductsInterface,
	dbInserter db.Inserter,
	w http.ResponseWriter,
	r *http.Request) {

	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		return
	}
	products, err := api.GetProducts(&http.Client{}, name)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	response, e := json.Marshal(products)
	if e != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(response)

	// Concurrent add products to database
	for _, v := range products {
		go func(prod db.Product) {
			dbInserter.Insert(prod, db.Database.Session, db.Name)
		}(v)
	}
}

// GetProductImageHandler func
func (p *ProductsAPI) GetProductImageHandler(
	api GetProductImageInterface,
	dbi db.FindInserter,
	w http.ResponseWriter,
	r *http.Request) {

	var (
		productID     = r.URL.Query().Get("product_id")
		side          = r.URL.Query().Get("side")
		dbQueryResult interface{}
		photo         db.Photo
		image         string
		ok            bool
		err           error
	)

	if len(productID) < 1 {
		w.WriteHeader(400)
		return
	}

	if len(side) < 1 {
		side = "100"
	}

	// Fetch from database
	dbQueryResult, err = dbi.Find(bson.M{"_id": productID, "side": side}, db.Photo{}, db.Database.Session, db.Name)
	if photo, ok = dbQueryResult.(db.Photo); ok == false {
		w.WriteHeader(500)
		log.Println("DbQueryResult return non photo value as expected")
		return
	}
	if len(photo.Data) < 10 || err != nil {
		if image, err = api.GetProductImageFromAPI(&http.Client{}, side, productID); err != nil {
			w.WriteHeader(500)
			return
		}

		response, e := json.Marshal(map[string]string{"result": image})
		if e != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(response)

		numProductID, e := strconv.Atoi(productID)

		if e == nil {
			numSide, e := strconv.Atoi(side)
			if e == nil {
				// Concurrent add photo to db
				go dbi.Insert(db.Photo{
					ProductID: numProductID,
					Side:      numSide,
					Data:      image,
				}, nil, nil)
			}
		}

	} else {
		response, e := json.Marshal(map[string]string{"result": photo.Data})
		if e != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(response)
	}

}

// GetProductShelflifeHandler func
func (p *ProductsAPI) GetProductShelflifeHandler(
	api GetProductShelfLifeInterface,
	dbi db.FindInserter,
	w http.ResponseWriter,
	r *http.Request) {

	var (
		productID     = r.URL.Query().Get("product_id")
		result        int
		response      []byte
		err           error
		dbQueryResult interface{}
		ok            bool
		shelf         db.Shelflife
	)

	if len(productID) < 1 {
		w.WriteHeader(400)
		return
	}

	if dbQueryResult, err = dbi.Find(bson.M{"_id": productID}, db.Shelflife{}, nil, nil); err != nil {
		var product db.Product
		err = nil
		if dbQueryResult, err = dbi.Find(bson.M{"_id": productID}, db.Product{}, nil, nil); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if product, ok = dbQueryResult.(db.Product); !ok {
			w.WriteHeader(500)
			return
		}

		if result, err = api.GetProductShelfLife(&http.Client{}, product.URL); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if response, err = json.Marshal(map[string]int{"result": result}); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if _, err = w.Write(response); err != nil {
			log.Println(err)
			return
		}

		go dbi.InsertShelflifeToDB(db.Shelflife{
			ProductID: product.ProductID,
			Data:      result,
		})
	} else {
		if response, err = json.Marshal(map[string]int{"result": shelf.Data}); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if _, err = w.Write(response); err != nil {
			log.Println(err)
			return
		}
	}
}
