package app

import (
	"context"
	"testing"

	"example.com/go-api/model"
)

type mockExtractor struct{}

func (m mockExtractor) Extract(ctx context.Context, num int) ([]model.Person, error) {
	return []model.Person{}, nil
}

type mockTransformer struct{}

func (m mockTransformer) Transform(input []model.Person) (map[string][]model.Person, error) {
	return nil, nil
}

type mockLoader struct{}

func (m mockLoader) Load(n map[string][]model.Person) error {
	return nil
}

func TestRun(t *testing.T) {
	err := Run(context.Background(), 10, mockExtractor{}, mockTransformer{}, mockLoader{})
	if err != nil {
		t.Errorf("expected test to pass, but it didn't. err:%v", err)
	}

}
