package server

import (
	"log"
	"net/http"

	"github.com/goncharovnikita/gosplitter"
)

// Server provides serving functions fy given port
type Server struct {
	Port string
	mux  *http.ServeMux
}

// Serve serve http server
func (s *Server) Serve() {
	log.Printf("Starting HTTP server on :%v", s.Port)
	s.mux = http.NewServeMux()
	var apiHandler = APIHandler{
		apiInterface: apiInterface,
		mux:          s.mux,
	}
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		w.Write([]byte("Hello!"))
	})
	gosplitter.Match("/api", s.mux, apiHandler)
	apiHandler.Start()
	log.Fatal(http.ListenAndServe(s.Port, s.mux))
}

// APIHandler handler
type APIHandler struct {
	mux *http.ServeMux
}

// GetProductsHandler handler
func (a *APIHandler) GetProductsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Add("Content-Type", "application/json")
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductsHandler(a.apiInterface, a.dbi, w, r)
	}
}

// GetProductImageHandler handler
func (a *APIHandler) GetProductImageHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Add("Content-Type", "application/json")
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductImageHandler(a.apiInterface, a.dbi, w, r)
	}
}

// GetProductShelflifeHandler handler
func (a *APIHandler) GetProductShelflifeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Add("Content-Type", "application/json")
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
