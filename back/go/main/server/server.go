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
	DB   *Database
}

// Serve serve http server
func (s *Server) Serve(apiInterface APIInterface) {
	log.Printf("Starting HTTP server on :%v", s.Port)
	s.mux = http.NewServeMux()
	var apiHandler = APIHandler{
		apiInterface: apiInterface,
		mux:          s.mux,
		dbi:          s.DB,
	}
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL.Path)
		w.Write([]byte("Hello!"))
	})
	gosplitter.Match("/api", s.mux, apiHandler)
	apiHandler.Start()
	log.Fatal(http.ListenAndServe(s.Port, s.mux))
}
