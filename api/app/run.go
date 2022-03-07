package app

import (
	"fmt"
	"log"

	"example.com/go-api/api"
	"example.com/go-api/datautil"
)

var num int

// Run is starting the execution for our application. Extract records -> Eliminate duplicates using set -> Grouping the records using a map -> Loading the groups in json files
func Run() error {
	// TODO: var name collides with imported package name, so it's advised to use a different name
	api := api.New("https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole")
	fmt.Println("Enter number of records:")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		return err
	}
	people, err := api.Extract(num)
	if err != nil {
		// TODO: Fatal will exit/break your script
		// Fatal is equivalent to Println() followed by a call to os.Exit(1).
		//func Fatalln(v ...interface{}) {
		//std.Output(2, fmt.Sprintln(v...))
		//os.Exit(1)
		//}
		log.Fatalf("err: %v", err)
		return err
	}
	groupedPeople, err := datautil.Transform(people)
	if err != nil {
		log.Fatalf("err:%v", err)
		return err
	}

	// TODO? is this useful?
	datautil.ShowMap(groupedPeople)

	err = datautil.Load(groupedPeople)
	if err != nil {
		log.Fatalf("err:%v", err)
		return err
	}
	return nil
}
