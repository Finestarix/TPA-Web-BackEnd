package hotel

import "github.com/graphql-go/graphql"

var hotelImageType *graphql.Object

func GetHotelImageType() *graphql.Object {
	if hotelImageType == nil {
		hotelImageType = graphql.NewObject(graphql.ObjectConfig{
			Name: "HotelImageType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"hotelid": &graphql.Field{
					Type: graphql.Int,
				},
				"source": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Hotel Image",
		})
	}
	return hotelImageType
}
