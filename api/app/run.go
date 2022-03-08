package app

import (
	"context"

	"example.com/go-api/etl"
)

// Run is starting the execution for our application. Extract records -> Eliminate duplicates using set -> Grouping the records using a map -> Loading the groups in json files
func Run(ctx context.Context, num int, e etl.Extractor, t etl.Transformer, l etl.Loader) error {

	people, err := e.Extract(ctx, num)
	if err != nil {
		return err
	}
	groupedPeople, err := t.Transform(people)
	if err != nil {
		return err
	}

	err = l.Load(groupedPeople)
	if err != nil {
		return err
	}
	return nil
}
