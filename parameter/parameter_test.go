package parameter_test

import (
	"bytes"
	"github.com/LeoCBS/garden/parameter"
	"io"
	"io/ioutil"
	"testing"
)

//TODO test body close

type readCloser struct {
	isClosed bool
	io.Reader
}

func (rc *readCloser) Close() error {
	rc.isClosed = true
	return nil
}

func setUp(rawJson string) *readCloser {
	return &readCloser{
		isClosed:  false,
		io.Reader: string.NewReader(rawJson),
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
}
