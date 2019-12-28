package main

import (
	"./graphql/queries"
	"./middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        queries.GetUserRoot(),
		Mutation:     nil,
	})

	if err != nil {
		panic(err.Error())
	}

	handlerSchema := handler.New(&handler.Config{
		Schema:           &schema,
		Pretty:           true,
		GraphiQL:         true,
		Playground:       true,
	})

	newHandlerSchema := middleware.SetCORS(handlerSchema)

	log.Fatal(http.ListenAndServe(":8080", newHandlerSchema))

}