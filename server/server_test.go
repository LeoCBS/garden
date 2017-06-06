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

func (m *mock) Save(body io.ReadCloser) (string, error) {
	if m.err {
		return "", errors.New("Save returned error")
	}
	return m.location, nil
}

func (m *mock) List() ([]byte, error) {
	if m.err {
		return nil, errors.New("List exploded!!")
	}
	return nil, nil
}

func TestSaveParameterSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/garden/v1/parameter/save", nil)
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

func TestSaveParameterError(t *testing.T) {
	req, err := http.NewRequest("POST", "/garden/v1/parameter/save", nil)
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

type fixture struct {
	req *http.Request
	rr  *httptest.ResponseRecorder
}

func setUp() {
}

func TestListParameterError(t *testing.T) {
	req, err := http.NewRequest("POST", "/garden/v1/parameter/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	s := server.NewServer(&mock{
		location: "",
		err:      true,
	})
	s.ServeMux.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
