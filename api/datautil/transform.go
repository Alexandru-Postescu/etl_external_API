package datautil

import (
	"fmt"

	"example.com/go-api/model"
)

// Transform loads the records of people into a set so we eliminate the duplicates and then we group the records based on the letter of the first name
// TODO: where is the error here? it always return nil
func Transform(input []model.Person) (map[string][]model.Person, error) {
	set := createSet(input)
	// TODO: don't use size on a map creation if it don't help you.
	// from the docs:
	//	Map: An empty map is allocated with enough space to hold the
	//	specified number of elements. The size may be omitted, in which case
	//	a small starting size is allocated.
	result := make(map[string][]model.Person)

	for person := range set.container {
		k := person.GetKey()
		result[k] = append(result[k], person)
	}
	return result, nil
}
func ShowMap(m map[string][]model.Person) {
	for k, v := range m {
		fmt.Printf("key[%v] value[%v]\n", k, v)
	}
}

// createSet creates a set of people using the records we extracted
func createSet(p []model.Person) customSet {
	set := makeSet()
	for _, person := range p {
		set.Add(person)
	}
	return *set
}
