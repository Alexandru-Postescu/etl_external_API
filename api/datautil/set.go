package datautil

import (
	"example.com/go-api/model"
)

type customSet map[model.Person]struct{}

//MakeSet initialize the set
// verify if it exists
func makeSet() customSet {
	return customSet{}
}

func (c customSet) Exists(key model.Person) bool {
	_, exists := c[key]
	return exists
}

func (c customSet) Add(key model.Person) {
	c[key] = struct{}{}
}
