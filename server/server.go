package server

import (
	"io"
	"net/http"
)

const (
	v1            = "/garden/v1"
	parameterPath = "/parameter"
)

type Storer interface {
	Store(io.ReadCloser) string
}

type Server struct {
	ServeMux *http.ServeMux
	st       Storer
}

func (s *Server) postParameterHandler(w http.ResponseWriter, r *http.Request) {
	location := s.st.Store(r.Body)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `{"alive": true}`)
}

func NewServer(str Storer) *Server {
	sm := http.NewServeMux()
	s := &Server{
		ServeMux: sm,
		st:       str,
	}
	sm.HandleFunc(v1+parameterPath, s.postParameterHandler)
	return s
}
