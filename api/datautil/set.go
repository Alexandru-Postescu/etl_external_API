package datautil

import (
	"fmt"

	"example.com/go-api/model"
)
// TODO: you could define a custom type directly on the map itself
// example: type customSet map[model.Person]struct{}, and add methods
// to the type(value receivers in this case, as maps are already reference types)
type customSet struct {
	container map[model.Person]struct{}
}

//MakeSet initialize the set
// verify if it exists
func makeSet() *customSet {
	return &customSet{
		container: make(map[model.Person]struct{}),
	}
}

func (c customSet) Exists(key model.Person) bool {
	_, exists := c.container[key]
	return exists
}

func (c *customSet) Add(key model.Person) {
	c.container[key] = struct{}{}
}

// TODO: this is not used
func (c *customSet) Remove(key model.Person) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}
