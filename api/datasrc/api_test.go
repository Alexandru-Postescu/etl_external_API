package datasrc

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-api/model"
)

func Test_Fetch(t *testing.T) {
	api := api{
		&http.Client{},
		"https://www.ardanlabs.com/blog/index.xml",
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

func Test_Extract(t *testing.T) {
	person1 := model.Person{
		"alex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	person2 := model.Person{
		"alexX",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	want := []model.Person{person1, person2}

	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(handleRequest))
	mockExtractor := NewAPI(ts.URL)

	got, err := mockExtractor.Extract(ctx, 2)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if len(want) != len(got) {
		t.Error("Number of records should be equal")
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	person1 := model.Person{
		"alex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	person2 := model.Person{
		"alexX",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	input := []model.Person{person1, person2}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
