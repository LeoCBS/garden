package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeoCBS/garden/server"
)

type mock struct {
	location string
}

//TODO test if put return error
func (m *mock) Post(body io.ReadCloser) (string, error) {
	return m.location, nil
}

func TestPostParameterHandler(t *testing.T) {
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
