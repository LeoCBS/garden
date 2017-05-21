package parameter_test

import (
	"bytes"
	"github.com/LeoCBS/garden/parameter"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

//TODO test body close

type readCloser struct {
	isClosed bool
	reader   io.Reader
}

func (rc *readCloser) Read(p []byte) (int, error) {
	return rc.reader.Read(p)
}

func (rc *readCloser) Close() error {
	rc.isClosed = true
	return nil
}

func setUp(rawJson string) *readCloser {
	return &readCloser{
		isClosed: false,
		reader:   strings.NewReader(rawJson),
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
		param := parameter.NewParameter()
		_, err := param.Store(value)
		if err == nil {
			t.Errorf("don't validate field correctly in test case %s", test)
		}
	}
}

func TestParameterFieldValid(t *testing.T) {
	rc := setUp(`{"name": "humidity", "value":1.0, "measure":"percent"}`)
	param := parameter.NewParameter()
	_, err := param.Store(rc)
	if err != nil {
		t.Error("Store return error with valid ReadCloser")
	}

}
