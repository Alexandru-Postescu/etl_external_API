package api

import (
	"net/http"
	"testing"
)
// TODO: testing with real data is not recommended!!! (imagine testing a credit card charge and using your approach)
func Test_Do(t *testing.T) {
	api := API{
		&http.Client{},
		"https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole",
	}
	request, err := api.Get()
	if err != nil {
		t.Error(err)
	}
	result, err := api.Do(request)
	if result.StatusCode > 299 {
		t.Errorf("Status code error: %v", result.StatusCode)
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Body) < 1 {
		t.Error("Didnt get a body")
	}

}

func TestAPI_Extract(t *testing.T) {
	api := API{
		&http.Client{},
		"https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole",
	}
	result, err := api.Extract(10)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 10 {
		t.Error("Didn't extract enough")
	}
}
