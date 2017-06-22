package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	v1                 = "/garden/v1"
	saveParameterPath  = "/parameter/save"
	listParametersPath = "/parameter/list"
)

type Parameter interface {
	Save(io.ReadCloser) (string, error)
	List() (interface{}, error)
}

type Server struct {
	ServeMux *http.ServeMux
	param    Parameter
	log      *log.Logger
}

// TODO check http method

func (s *Server) listParametersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		err := errors.New(http.StatusText(http.StatusMethodNotAllowed))
		s.errorHandler(w, err, http.StatusMethodNotAllowed)
		return
	}

	parameters, err := s.param.List()
	if err != nil {
		s.errorHandler(w, err, http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(&parameters)
	if err != nil {
		s.errorHandler(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) errorHandler(w http.ResponseWriter, err error, errorCode int) {
	s.log.Println(err)
	w.WriteHeader(errorCode)
	io.WriteString(w, err.Error())
}

func (s *Server) saveParameterHandler(w http.ResponseWriter, r *http.Request) {
	location, err := s.param.Save(r.Body)
	if err != nil {
		s.errorHandler(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `created`)
}

func NewServer(p Parameter) *Server {
	sm := http.NewServeMux()
	info := log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	s := &Server{
		ServeMux: sm,
		param:    p,
		log:      info,
	}
	sm.HandleFunc(v1+saveParameterPath, s.saveParameterHandler)
	sm.HandleFunc(v1+listParametersPath, s.listParametersHandler)
	return s
}
