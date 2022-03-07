package main

import (
	//"context"
	"context"
	"log"

	"example.com/go-api/app"
)

func main() {
	ctx := context.Background()
	err := app.Run(ctx)
	if err != nil {
		log.Fatalf("err:%v", err)
	}

}
