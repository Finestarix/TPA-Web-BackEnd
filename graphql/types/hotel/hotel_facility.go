package hotel

import "github.com/graphql-go/graphql"

var hotelFacilityType *graphql.Object

func GetHotelFacilityType() *graphql.Object {
	if hotelFacilityType == nil {
		hotelFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name: "HotelFacilityType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"hotelid": &graphql.Field{
					Type: graphql.Int,
				},
				"photo": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Hotel Facility",
		})
	}
	return hotelFacilityType
}