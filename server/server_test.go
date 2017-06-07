package server_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeoCBS/garden/server"
)

type ParamJson struct {
	Name string `json:"name"`
}

type mock struct {
	location   string
	err        bool
	parameters ParamJson
}

func (m *mock) Save(body io.ReadCloser) (string, error) {
	if m.err {
		return "", errors.New("Save returned error")
	}
	return m.location, nil
}

func (m *mock) List() (interface{}, error) {
	if m.err {
		return nil, errors.New("List exploded!!")
	}
	return m.parameters, nil
}

type fixture struct {
	req *http.Request
	rr  *httptest.ResponseRecorder
	s   *server.Server
}

func setUp(
	t *testing.T,
	httpMethod string,
	path string,
	isError bool,
	expecParams ParamJson,
	expecLocation string,
) *fixture {
	req := httptest.NewRequest(httpMethod, path, nil)
	rr := httptest.NewRecorder()
	s := server.NewServer(&mock{
		location:   expecLocation,
		err:        isError,
		parameters: expecParams,
	})
	s.ServeMux.ServeHTTP(rr, req)
	return &fixture{
		req: req,
		rr:  rr,
		s:   s,
	}
}

func TestSaveParameterSuccess(t *testing.T) {
	expectedLocation := "stored"
	f := setUp(t, "POST", "/garden/v1/parameter/save", false, ParamJson{}, expectedLocation)

	if status := f.rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	resp := f.rr.Result()
	location := resp.Header.Get("Location")
	if location != expectedLocation {
		t.Error("server don't return expected location")
	}
}

func TestSaveParameterError(t *testing.T) {
	f := setUp(t, "POST", "/garden/v1/parameter/save", true, ParamJson{}, "")

	if status := f.rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TODO test invalid http method

func TestListParameterError(t *testing.T) {
	f := setUp(t, "GET", "/garden/v1/parameter/list", true, ParamJson{}, "")

	if status := f.rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestListParameterSuccess(t *testing.T) {
	expectedParam := ParamJson{Name: "ok"}
	f := setUp(t, "GET", "/garden/v1/parameter/list", false, expectedParam, "")

	if status := f.rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	resp := f.rr.Result()
	expectedContentType := "application/json"
	receivedContentType := resp.Header.Get("Content-Type")
	if expectedContentType != receivedContentType {
		t.Error("server response wrong content type, expected = %s, received = %s",
			expectedContentType, receivedContentType)
	}
	receivedParam := ParamJson{}
	json.NewDecoder(resp.Body).Decode(&receivedParam)
	if expectedParam.Name != receivedParam.Name {
		t.Errorf("List don't return expected body, expected = %s, received = %s",
			expectedContentType, receivedContentType)
	}
}
