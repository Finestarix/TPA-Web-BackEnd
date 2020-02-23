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
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"locationid": &graphql.Field{
					Type: graphql.Int,
				},
				"photo": {
					Type: graphql.NewList(GetHotelImageType()),
				},
				"facility": {
					Type: graphql.NewList(GetHotelFacilityType()),
				},
				"type": {
					Type: graphql.NewList(GetHotelTypeType()),
				},
				"latitude": &graphql.Field{
					Type: graphql.Float,
				},
				"longitude": &graphql.Field{
					Type: graphql.Float,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"rating": &graphql.Field{
					Type: graphql.Float,
				},
				"information": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Hotel",
		})
	}
	return hotelType
}
