package main

import (
	"context"
	"log"

	"example.com/go-api/app"
)

func main() {
	ctx := context.Background()
	err := app.Run()
	if err != nil {
		log.Fatalf("err:%v", err)
	}

}
