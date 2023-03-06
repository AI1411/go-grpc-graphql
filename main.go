package main

import (
	"fmt"
	"log"

	"github.com/AI1411/go-graphql-grpc/internal/env"
	"github.com/AI1411/go-graphql-grpc/internal/infra/db"
)

func main() {
	e, err := env.NewValue()
	if err != nil {
		fmt.Println(err.Error())
		panic("Error loading .env file")
	}
	log.Printf("env=%+v", e)

	dbClient, err := db.NewClient(&e.DB)
	if err != nil {
		panic(err)
	}
	log.Printf("dbClient=%+v", dbClient)
}
