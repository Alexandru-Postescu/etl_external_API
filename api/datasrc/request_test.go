package datasrc

import (
	"context"
	"net/http"
	"testing"
)

// TODO: testing with real data is not recommended!!! (imagine testing a credit card charge and using your approach)
func Test_Fetch(t *testing.T) {
	api := API{
		&http.Client{},
		"https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole",
	}
	ctx := context.Background()

	result, err := api.FetchData(ctx)
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
	ctx := context.Background()
	result, err := api.Extract(ctx, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 10 {
		t.Error("Didn't extract enough")
	}
}
