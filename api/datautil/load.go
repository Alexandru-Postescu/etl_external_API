package datautil

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"example.com/go-api/model"
)

// Mapping the map to the required output
func mapOutput(m map[string][]model.Person) []model.Output {
	// TODO: using make like this, brings no value, instead it creates an extra allocation!!!
	output := make([]model.Output, 0)
	for k, v := range m {
		o := model.Output{}
		o.Init(k, v, len(v))
		output = append(output, o)
	}
	return output
}

// Loading the groups of people in json files
func Load(m map[string][]model.Person) error {
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
