package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"example.com/go-api/model"
)

// TODO: package api, object API (api.API)? sounds like shit :)
// I also think you can find a better name
// API object is used to handle requests and to extract the number of records read from the keyboard
type API struct {
	Client *http.Client
	URL    string
}

// TODO: if you choose value semantics, stick to that and return a value here, not a pointer
func New(url string) *API {
	a := new(API)
	a.URL = url
	a.Client = http.DefaultClient
	return a
}

// Extract returns the records of people
func (a API) Extract(num int) ([]model.Person, error) {

	if num < 1 && num > 10000 {
		// TODO: based on what?
		return nil, errors.New("number is too low or too high")
	}
	p := make([]model.Person, 0, num)
	request, err := a.Get()
	if err != nil {
		return p, err
	}

	for num > 0 {
		response, err := a.Do(request)
		if err != nil {
			return p, err
		}
		if response.StatusCode > 299 {
			return p, errors.New("status code error")
		}
		var temp []model.Person
		result := response.Body
		err = json.Unmarshal(result, &temp)
		if err != nil {
			return nil, err
		}

		if num < len(temp) {
			p = append(p, temp[0:num]...)
			break
		}

		p = append(p, temp...)
		num = num - len(temp)

	}
	print(len(p))
	return p, nil

}

// TODO: confusing naming
// Get is creating the request
func (a API) Get() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, a.URL, nil)
}

// Do is executing the request
func (a API) Do(request *http.Request) (ResponseStruct, error) {

	r := ResponseStruct{}
	response, err := a.Client.Do(request)
	if err != nil {
		return r, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r, err
	}

	// TODO: body can be closed in a defer without checking the error
	err = response.Body.Close()
	if err != nil {
		return r, err
	}

	return ResponseStruct{
		StatusCode: response.StatusCode,
		Body:       body,
	}, nil
}
