package datautil

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"example.com/go-api/etl"
	"example.com/go-api/model"
)

// This struct implements the load interface, using it to load the grouped people into json files
type jsonLoader struct{}

func NewJSONLoader() etl.Loader {
	return jsonLoader{}
}

// Loading the groups of people in json files
func (j jsonLoader) Load(m map[string][]model.Person) error {
	o := mapOutput(m)

	count := 0
	for _, p := range o {
		name := "file" + strconv.Itoa(count) + ".json"
		file, err := json.MarshalIndent(p, "", " ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(name, file, 0644)
		if err != nil {
			return err
		}
		count++
	}
	return nil
}

// Mapping the map to the required output
func mapOutput(m map[string][]model.Person) []model.Output {
	var output []model.Output
	for k, v := range m {
		o := model.Output{}
		o.Init(k, v, len(v))
		output = append(output, o)
	}
	return output
}
