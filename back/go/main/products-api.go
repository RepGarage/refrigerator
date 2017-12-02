package main

import "net/http"
import "encoding/json"

// ProductsAPI type
type ProductsAPI struct{}

var apiInterface ApiInterface

// GetProductsHandler func
func (p *ProductsAPI) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Add("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		return
	}
	products, err := apiInterface.GetProducts(name)
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
}

// GetProductImageHandler func
func (p *ProductsAPI) GetProductImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Add("Content-Type", "application/json")
	productID := r.URL.Query().Get("product_id")
	side := r.URL.Query().Get("side")
	if len(productID) < 1 {
		w.WriteHeader(400)
		return
	}

	if len(side) < 1 {
		side = "100"
	}

	image, err := apiInterface.GetProductImageFromAPI(side, productID)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	response, e := json.Marshal(map[string]string{"result": image})
	if e != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(response)
}
