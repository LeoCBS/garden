package server

import (
	"fmt"
	"io"
	"net/http"
)

const (
	v1            = "/garden/v1"
	parameterPath = "/parameter"
)

type Parameter interface {
	Put(io.ReadCloser) (string, error)
}

type Server struct {
	ServeMux *http.ServeMux
	param    Parameter
}

// TODO implements handler erros (501,  404)

func (s *Server) postParameterHandler(w http.ResponseWriter, r *http.Request) {
	location, err := s.param.Put(r.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Internal error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, errMsg)
		return
	}

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `created`)
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
