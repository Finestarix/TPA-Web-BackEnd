package event

import "github.com/graphql-go/graphql"

var eventType *graphql.Object

func GetEntertainmentType() *graphql.Object {
	if eventType == nil {
		eventType = graphql.NewObject(graphql.ObjectConfig{
			Name: "EventType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"location": &graphql.Field{
					Type: graphql.String,
				},
				"latitude": &graphql.Field{
					Type: graphql.Float,
				},
				"longitude": &graphql.Field{
					Type: graphql.Float,
				},
				"date": &graphql.Field{
					Type: graphql.String,
				},
				"category": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Blog Model",
		})
	}
	return eventType
}


