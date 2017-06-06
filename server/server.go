package server

import (
	"fmt"
	"io"
	"net/http"
)

const (
	v1                 = "/garden/v1"
	saveParameterPath  = "/parameter/save"
	listParametersPath = "/parameter/list"
)

type Parameter interface {
	Save(io.ReadCloser) (string, error)
	List() ([]byte, error)
}

type Server struct {
	ServeMux *http.ServeMux
	param    Parameter
}

// TODO implements handler erros (501,  404)

func (s *Server) listParametersHandler(w http.ResponseWriter, r *http.Request) {
	parameters, err := s.param.List()
	if err != nil {
		errorHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(parameters)
}

func errorHandler(w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("Internal error: %s", err)
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, errMsg)
}

func (s *Server) saveParameterHandler(w http.ResponseWriter, r *http.Request) {
	location, err := s.param.Save(r.Body)
	if err != nil {
		errorHandler(w, err)
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
	sm.HandleFunc(v1+saveParameterPath, s.saveParameterHandler)
	sm.HandleFunc(v1+listParametersPath, s.listParametersHandler)
	return s
}
