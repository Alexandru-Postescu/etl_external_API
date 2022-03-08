package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"example.com/go-api/app"
	"example.com/go-api/datasrc"
	"example.com/go-api/datautil"
)

var num int
var errNum error

func init() {
	fmt.Println("Enter number of records:")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		errNum = errors.New("bad argument")
	}
}
func main() {
	if errNum != nil {
		log.Fatalf("err: %v", errNum)
	}
	if num < 0 {
		log.Fatalf("Number of records should be more than 1")
	}

	ctx := context.Background()
	extractor := datasrc.NewAPI("https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole")
	tranformer := datautil.NewDataTransformer()
	loader := datautil.NewJSONLoader()

	err := app.Run(ctx, num, extractor, tranformer, loader)
	if err != nil {
		log.Fatalf("err:%v", err)
	}

}
