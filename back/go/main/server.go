package main

import (
	"log"
	"net/http"

	"github.com/goncharovnikita/gosplitter"
)

// Server provides serving functions fy given port
type Server struct {
	port string
	mux  *http.ServeMux
	db   *Database
}

// Serve serve http server
func (s *Server) Serve(apiInterface APIInterface) {
	log.Printf("Starting HTTP server on :%v", s.port)
	s.mux = http.NewServeMux()
	var apiHandler = APIHandler{
		apiInterface: apiInterface,
		mux:          s.mux,
		dbi:          s.db,
	}
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		w.Write([]byte("Hello!"))
	})
	gosplitter.Match("/api", s.mux, apiHandler)
	apiHandler.Start()
	log.Fatal(http.ListenAndServe(s.port, s.mux))
}
