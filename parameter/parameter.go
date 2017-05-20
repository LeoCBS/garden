package server

import (
	"encoding/json"
	"io"
)

type Parameter struct{}

type ParameterJson struct {
	Name    string `json:"name"`
	Value   float  `json:"value"`
	Measure string `json:"measure"`
}

func (s *Parameter) Store(body io.ReadCloser) (string, error) {
	param, err := decodeJson(body)
	if err != nil {
		return "", err
	}
	err = validateParameterFields(param)
	if err != nil {
		return "", err
	}
}

func validateParameterFields(param Parameter) error {
	if param.Name == "" {
		return errors.New("field name is required")
	}
	if param.Value == "" {
		return errors.New("field value is required")
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
		return "", err
	}
	defer body.Close()
	return param
}

func NewParameter() *Parameter {
	return &Parameter{}
}
