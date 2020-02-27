package flight

import (
	"github.com/graphql-go/graphql"
)

var flightCompanyType *graphql.Object

func GetFlightCompanyType() *graphql.Object {
	if flightCompanyType == nil {
		flightCompanyType = graphql.NewObject(graphql.ObjectConfig{
			Name: "FlightCompanyType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Flight Company",
		})
	}
	return flightCompanyType
}
