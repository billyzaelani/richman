package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server is a richman http server.
type Server struct {
	r *mux.Router
}

// New returns a new server.
func New(options ...Option) *Server {
	srv := &Server{
		r: mux.NewRouter(),
	}

	for _, option := range options {
		option(srv)
	}

	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
