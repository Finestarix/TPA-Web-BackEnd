package types

import "github.com/graphql-go/graphql"

var carModelType *graphql.Object

func GetCarModelType() *graphql.Object {
	if carModelType == nil {
		carModelType = graphql.NewObject(graphql.ObjectConfig{
			Name: "CarModel",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"brand": &graphql.Field{
					Type: graphql.String,
				},
				"model": &graphql.Field{
					Type: graphql.String,
				},
				"passenger": &graphql.Field{
					Type: graphql.Int,
				},
				"baggage": &graphql.Field{
					Type: graphql.Int,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Car Model",
		})
	}
	return carModelType
}
