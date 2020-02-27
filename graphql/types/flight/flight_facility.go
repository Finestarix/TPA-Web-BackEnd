package flight

import (
	"github.com/graphql-go/graphql"
)

var flightFacilityType *graphql.Object

func GetFlightFacilityType() *graphql.Object {
	if flightFacilityType == nil {
		flightFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name: "FlightFacilityType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"flightID": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "Type for Flight Facility",
		})
	}
	return flightFacilityType
}
