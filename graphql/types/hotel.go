package types

import "github.com/graphql-go/graphql"

var hotelType *graphql.Object

func GetHotelType() *graphql.Object {
	if hotelType == nil {
		hotelType = graphql.NewObject(graphql.ObjectConfig{
			Name: "HotelType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"rating": &graphql.Field{
					Type: graphql.Float,
				},
			},
			Description: "Type for Hotel",
		})
	}
	return hotelType
}
