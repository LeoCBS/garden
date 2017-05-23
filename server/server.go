package server

import (
	"io"
	"net/http"
)

const (
	v1            = "/garden/v1"
	parameterPath = "/parameter"
)

type Parameter interface {
	Post(io.ReadCloser) (string, error)
}

type Server struct {
	ServeMux *http.ServeMux
	param    Parameter
}

func (s *Server) postParameterHandler(w http.ResponseWriter, r *http.Request) {
	location, _ := s.param.Post(r.Body)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `{"alive": true}`)
}

func NewServer(p Parameter) *Server {
	sm := http.NewServeMux()
	s := &Server{
		ServeMux: sm,
		param:    p,
	}
	sm.HandleFunc(v1+parameterPath, s.postParameterHandler)
	return s
}
