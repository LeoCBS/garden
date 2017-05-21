package parameter

import (
	"encoding/json"
	"errors"
	"io"
)

type Storer interface {
	Store(interface{}) error
}

type Parameter struct {
	s Storer
}

type ParameterJson struct {
	Name    string  `json:"name"`
	Value   float32 `json:"value"`
	Measure string  `json:"measure"`
}

func (param *Parameter) Put(body io.ReadCloser) (string, error) {
	paramJson, err := decodeJson(body)
	if err != nil {
		return "", err
	}
	err = validateParameterFields(paramJson)
	if err != nil {
		return "", err
	}
	err = param.s.Store(paramJson)
	if err != nil {
		return "", err
	}
	return "", err
}

func validateParameterFields(param ParameterJson) error {
	if param.Name == "" {
		return errors.New("field name is required")
	}
	if param.Measure == "" {
		return errors.New("field measure is required")
	}
	return nil
}

func decodeJson(body io.ReadCloser) (ParameterJson, error) {
	d := json.NewDecoder(body)
	var param ParameterJson
	err := d.Decode(&param)
	if err != nil {
		return ParameterJson{}, err
	}
	defer body.Close()
	return param, err
}

func NewParameter(storer Storer) *Parameter {
	return &Parameter{
		s: storer,
	}
}
