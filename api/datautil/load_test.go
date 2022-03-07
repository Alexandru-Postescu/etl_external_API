package datautil

import (
	"testing"

	"example.com/go-api/model"
)

func TestLoad(t *testing.T) {

	person1 := model.Person{
		"alex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	person2 := model.Person{
		"aalex",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	input1 := []model.Person{person1, person2}
	input := map[string][]model.Person{
		"a": input1,
	}
	err := Load(input)
	if err != nil {
		t.Errorf("Load() error: %v", err)
	}

}
