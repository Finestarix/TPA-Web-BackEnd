package types

import "github.com/graphql-go/graphql"

var locationType *graphql.Object

func GetLocationType() *graphql.Object {
	if locationType == nil {
		locationType = graphql.NewObject(graphql.ObjectConfig{
			Name: "LocationType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"longitude": &graphql.Field{
					Type: graphql.Float,
				},
				"latitude": &graphql.Field{
					Type: graphql.Float,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"province": &graphql.Field{
					Type: graphql.String,
				},
				"region": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Location",
		})
	}
	return locationType
}
