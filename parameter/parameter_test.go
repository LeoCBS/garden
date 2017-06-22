package parameter_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/LeoCBS/garden/parameter"
)

type fixture struct {
	rc    *readCloser
	param *parameter.Parameter
}

type readCloser struct {
	isClosed    bool
	isReadError bool
	reader      io.Reader
}

func (rc *readCloser) Read(p []byte) (int, error) {
	if rc.isReadError {
		return 0, errors.New("Read throws error")
	}
	return rc.reader.Read(p)
}

func (rc *readCloser) Close() error {
	rc.isClosed = true
	return nil
}

type storerMock struct {
	isStoreError bool
	isLoadError  bool
	expecParams  interface{}
}

func (m *storerMock) Store(interface{}) error {
	if m.isStoreError {
		return errors.New("Store exploded")
	}
	return nil
}

func (m *storerMock) Load() (interface{}, error) {
	if m.isLoadError {
		return nil, errors.New("Load exploded!!")
	}
	return m.expecParams, nil
}

func setUp(
	isReadError bool,
	isStoreError bool,
	isLoadError bool,
	expecParams interface{},
) *fixture {
	validJson := `{"name": "humidity", "value":1.0, "measure":"percent"}`
	rc := &readCloser{
		isClosed:    false,
		isReadError: isReadError,
		reader:      strings.NewReader(validJson),
	}
	param := parameter.NewParameter(&storerMock{
		isStoreError: isStoreError,
		isLoadError:  isLoadError,
		expecParams:  expecParams,
	})
	return &fixture{
		rc:    rc,
		param: param,
	}
}

func TestParameterFieldsInvalids(t *testing.T) {
	testcases := map[string]io.ReadCloser{
		"nameEmpty": ioutil.NopCloser(
			bytes.NewReader([]byte(`{"value":80.0, "measure":"percent"}`))),
		"measureEmpty": ioutil.NopCloser(
			bytes.NewReader([]byte(`{"name":"test", "value":0.0}`))),
	}
	for test, value := range testcases {
		param := parameter.NewParameter(&storerMock{})
		_, err := param.Save(value)
		if err == nil {
			t.Errorf("don't validate field correctly in test case %s", test)
		}
	}
}

func TestNewParameterFieldValid(t *testing.T) {
	isReadError := false
	isStoreError := false
	isLoadError := false
	f := setUp(isReadError, isStoreError, isLoadError, nil)
	_, err := f.param.Save(f.rc)
	if err != nil {
		t.Error("Store return error with valid ReadCloser")
	}
}

func TestShouldCatchReadError(t *testing.T) {
	isReadError := true
	isStoreError := false
	isLoadError := false
	f := setUp(isReadError, isStoreError, isLoadError, nil)
	_, err := f.param.Save(f.rc)
	if err == nil {
		t.Error("Read don't return error as expected")
	}
}

func TestShouldReadCloseWithSuccess(t *testing.T) {
	isReadError := false
	isStoreError := false
	isLoadError := false
	f := setUp(isReadError, isStoreError, isLoadError, nil)
	_, err := f.param.Save(f.rc)
	if err != nil {
		t.Error("Store return error with valid ReadCloser")
	}
	if !f.rc.isClosed {
		t.Error("Store don't call Close func, possible leak memory")
	}
}

func TestShouldGetErrorOnStore(t *testing.T) {
	isReadError := false
	isStoreError := true
	isLoadError := false
	f := setUp(isReadError, isStoreError, isLoadError, nil)
	_, err := f.param.Save(f.rc)
	if err == nil {
		t.Error("Store() don't return error as expected")
	}
}

func TestShouldReadCloseAndStoreWithSuccess(t *testing.T) {
	isReadError := false
	isStoreError := false
	isLoadError := false
	f := setUp(isReadError, isStoreError, isLoadError, nil)
	_, err := f.param.Save(f.rc)
	if !f.rc.isClosed {
		t.Error("Store don't call Close func, possible leak memory")
	}
	if err != nil {
		t.Error("Store() don't return error as expected")
	}
}

func TestShouldGetErrorOnLoad(t *testing.T) {
	isReadError := false
	isStoreError := false
	isLoadError := true
	f := setUp(isReadError, isStoreError, isLoadError, nil)
	_, err := f.param.List()
	if err == nil {
		t.Error("Load() don't return error as expected")
	}
}

func TestShouldListWithSuccess(t *testing.T) {
	isReadError := false
	isStoreError := false
	isLoadError := false
	expecParams := `[{"name": "humidity", "value":1.0, "measure":"percent"},
			{"name": "humidity", "value":1.0, "measure":"percent"}]`
	f := setUp(isReadError, isStoreError, isLoadError, expecParams)
	params, err := f.param.List()
	if err != nil {
		t.Error("Load() don't return error as expected")
	}
	if params != expecParams {
		t.Error("Load() don't return expected value")
	}
}
