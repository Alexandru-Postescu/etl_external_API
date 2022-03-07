package app

import (
	"context"
	"fmt"

	"example.com/go-api/datasrc"
	"example.com/go-api/datautil"
)

var num int

// Run is starting the execution for our application. Extract records -> Eliminate duplicates using set -> Grouping the records using a map -> Loading the groups in json files
func Run(ctx context.Context) error {
	// TODO: var name collides with imported package name, so it's advised to use a different name
	apiSource := datasrc.NewAPI("https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole")
	fmt.Println("Enter number of records:")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		return err
	}
	people, err := apiSource.Extract(ctx, num)
	if err != nil {
		return err
	}
	groupedPeople, err := datautil.Transform(people)
	if err != nil {
		return err
	}

	err = datautil.Load(groupedPeople)
	if err != nil {
		return err
	}
	return nil
}
