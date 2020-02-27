package hotel

import "github.com/graphql-go/graphql"

var hotelTypeType *graphql.Object

func GetHotelTypeType() *graphql.Object {
	if hotelTypeType == nil {
		hotelTypeType = graphql.NewObject(graphql.ObjectConfig{
			Name: "HotelTypeType",
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
			Description: "Type for Hotel Type",
		})
	}
	return hotelTypeType
}