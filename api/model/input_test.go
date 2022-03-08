package model

import (
	"testing"
)

func TestPerson_GetKey(t *testing.T) {
	p := Person{
		"a",
		"poste",
		"idk",
		"idk",
		"idk",
		"idk",
	}
	letter, err := p.GetKey()
	if err != nil {
		t.Errorf("err:%v", err)
	}
	if letter != "a" {
		t.Error("Test failed.")
	}
}
