package main

import (
	"log"
	"net/http"
)

var productsAPI ProductsAPI

// Server provides serving functions fy given port
type Server struct {
	port string
	mux  *http.ServeMux
	db   *DatabaseInterface
}

// Serve serve http server
func (s *Server) Serve() {
	log.Printf("Starting HTTP server on :%v", s.port)
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		w.Write([]byte("Hello!"))
	})
	s.mux.HandleFunc("/api/get/products", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductsHandler(w, r)
	})
	s.mux.HandleFunc("/api/get/product/image", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		productsAPI.GetProductImageHandler(w, r)
	})
	log.Fatal(http.ListenAndServe(s.port, s.mux))
}
