package server

import (
	"log"
	"net/http"

	"github.com/goncharovnikita/gosplitter"
)

var productsAPI ProductsAPI

// APIHandler handler
type APIHandler struct {
	mux          *http.ServeMux
	dbi          DatabaseInterface
	apiInterface APIInterface
}

// GetProductsHandler handler
func (a *APIHandler) GetProductsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductsHandler(a.apiInterface, a.dbi, w, r)
	}
}

// GetProductImageHandler handler
func (a *APIHandler) GetProductImageHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductImageHandler(a.apiInterface, a.dbi, w, r)
	}
}

// GetProductShelflifeHandler handler
func (a *APIHandler) GetProductShelflifeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductShelflifeHandler(a.apiInterface, a.dbi, w, r)
	}
}

// Start serving
func (a *APIHandler) Start() {
	gosplitter.Match("/get/products", a.mux, a.GetProductsHandler())
	gosplitter.Match("/get/product/image", a.mux, a.GetProductImageHandler())
	gosplitter.Match("/get/product/shelflife", a.mux, a.GetProductShelflifeHandler())
}
