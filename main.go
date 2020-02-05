package main

import (
	"./graphql/mutations"
	"./graphql/queries"
	"./middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queries.Root(),
		Mutation: mutations.Root(),
	})

	if err != nil {
		panic(err.Error())
	}

	handlerSchema := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	newHandlerSchema := middleware.SetCORS(handlerSchema)

	log.Fatal(http.ListenAndServe(":4201", newHandlerSchema))
}

//func main() {
	//var temp = models.GetNearestHotel(-6.202112409236046, 106.78171499648295)
	//db := connection.GetConnection()
	//for i, _ := range temp {
	//	db.Model(temp[i]).Related(&temp[i].Location, "location_id")
	//	fmt.Println(temp[i])
	//}

	//models.GenerateLocation()
//}