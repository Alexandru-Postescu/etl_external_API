package api

import (
	"encoding/json"

	"example.com/go-api/model"
)

// ResponseStruct is used to save the responses from the requests
type ResponseStruct struct {
	StatusCode int
	Body       []byte
}

func (r ResponseStruct) Get() ResponseStruct {
	return r
}
func (r ResponseStruct) To(value []model.Person) error {
	err := json.Unmarshal(r.Body, &value)
	if err != nil {
		value = nil
		return err
	}
	return nil
}
