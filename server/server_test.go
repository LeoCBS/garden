package server_test

import (
	//	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
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
		location:   expecParams,
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

// TODO test invalid http method

func TestListParameterError(t *testing.T) {
	f := setUp(t, "GET", "/garden/v1/parameter/list", true, ParamJson{})

	if status := f.rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestListParameterSuccess(t *testing.T) {
	expectedParameters := ParamJson{Name: "ok"}
	t.Error(expectedParameters)
	f := setUp(t, "GET", "/garden/v1/parameter/list", false, expectedParameters)

	if status := f.rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	resp := f.rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Error(resp.Header.Get("Content-Type"))
	receivedParam := ParamJson{}
	json.NewDecoder(resp.Body).Decode(&receivedParam)
	//t.Error(receivedParam)
	//err := json.Unmarshal(body, &p)
	//if err != nil {
	//		t.Error(err)
	//	}
	//	t.Error(p)
	//if expectedParameters.Test != receivedParam.Test {
	//	t.Errorf("List don't return expected body")
	//}
}
