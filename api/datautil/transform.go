package datautil

import (
	"example.com/go-api/etl"
	"example.com/go-api/model"
)

// Transform loads the records of people into a set so we eliminate the duplicates and then we group the records based on the letter of the first name
type dataTransformer struct{}

func NewDataTransformer() etl.Transformer {
	return dataTransformer{}
}
func (d dataTransformer) Transform(input []model.Person) (map[string][]model.Person, error) {
	setPeople := createSetPeople(input)

	result := make(map[string][]model.Person)

	for person := range setPeople {
		k, err := person.GetKey()
		if err != nil {
			return nil, err
		}
		result[k] = append(result[k], person)
	}
	return result, nil
}

// createSet creates a set of people using the records we extracted
func createSetPeople(p []model.Person) customSet {
	set := makeSet()
	for _, person := range p {
		set.Add(person)
	}
	return set
}
