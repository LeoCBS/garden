package server_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeoCBS/garden/server"
)

type mock struct {
	location string
	err      bool
}

//TODO test if put return error
func (m *mock) Put(body io.ReadCloser) (string, error) {
	if m.err {
		return "", errors.New("Put returned error")
	}
	return m.location, nil
}

func TestPutParameterSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/garden/v1/parameter", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	expectedLocation := "stored"
	s := server.NewServer(&mock{
		location: expectedLocation,
	})
	s.ServeMux.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	resp := rr.Result()
	location := resp.Header.Get("Location")
	if location != expectedLocation {
		t.Error("server don't return expected location")
	}
}

func TestPutParameterError(t *testing.T) {
	req, err := http.NewRequest("POST", "/garden/v1/parameter", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	expectedLocation := "stored"
	s := server.NewServer(&mock{
		location: expectedLocation,
		err:      true,
	})
	s.ServeMux.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
