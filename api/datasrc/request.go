package datasrc

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"example.com/go-api/model"
)

type API struct {
	Client *http.Client
	URL    string
}

func NewAPI(url string) API {
	a := API{}
	a.URL = url
	a.Client = http.DefaultClient
	return a
}

// Extract returns the records of people
func (a API) Extract(ctx context.Context, num int) ([]model.Person, error) {

	if num < 1 {

		return nil, errors.New("number is too low")
	}
	p := make([]model.Person, 0, num)

	for num > 0 {
		response, err := a.FetchData(ctx)
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

	return p, nil
}

// Fetch data is making a get request to the API and returns its response (status code and body)
func (a API) FetchData(ctx context.Context) (ResponseStruct, error) {
	method := "GET"
	request, err := http.NewRequestWithContext(ctx, method, a.URL, nil)
	r := ResponseStruct{}
	if err != nil {
		return r, err
	}

	response, err := a.Client.Do(request)
	if err != nil {
		return r, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r, err
	}

	defer response.Body.Close()

	return ResponseStruct{
		StatusCode: response.StatusCode,
		Body:       body,
	}, nil
}
