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
	isError bool
}

func (m *storerMock) Store(interface{}) error {
	if m.isError {
		return errors.New("Store exploded")
	}
	return nil
}

func setUp(isReadError bool, isStoreError bool) *fixture {
	validJson := `{"name": "humidity", "value":1.0, "measure":"percent"}`
	rc := &readCloser{
		isClosed:    false,
		isReadError: isReadError,
		reader:      strings.NewReader(validJson),
	}
	param := parameter.NewParameter(&storerMock{
		isError: isStoreError,
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
	f := setUp(isReadError, isStoreError)
	_, err := f.param.Save(f.rc)
	if err != nil {
		t.Error("Store return error with valid ReadCloser")
	}
}

func TestShouldCatchReadError(t *testing.T) {
	isReadError := true
	isStoreError := false
	f := setUp(isReadError, isStoreError)
	_, err := f.param.Save(f.rc)
	if err == nil {
		t.Error("Read don't return error as expected")
	}
}

func TestShouldReadCloseWithSuccess(t *testing.T) {
	isReadError := false
	isStoreError := false
	f := setUp(isReadError, isStoreError)
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
	f := setUp(isReadError, isStoreError)
	_, err := f.param.Save(f.rc)
	if err == nil {
		t.Error("Store() don't return error as expected")
	}
}

func TestShouldReadCloseAndStoreWithSuccess(t *testing.T) {
	isReadError := false
	isStoreError := false
	f := setUp(isReadError, isStoreError)
	_, err := f.param.Save(f.rc)
	if !f.rc.isClosed {
		t.Error("Store don't call Close func, possible leak memory")
	}
	if err != nil {
		t.Error("Store() don't return error as expected")
	}
}
