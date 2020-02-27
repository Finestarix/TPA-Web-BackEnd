package flight

import (
	models "../../models/flight"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllFlightFacility(p graphql.ResolveParams) (i interface{}, err error) {
	flightFacility := models.GetAllFlightFacility()
	if len(flightFacility) == 0 {
		return nil, errors.New("error: no data found")
	}

	return flightFacility, nil
}

func InsertFlightFacility(p graphql.ResolveParams) (i interface{}, err error) {
	flightID := p.Args["flightID"].(int)
	name := p.Args["name"].(string)

	newFlightFacility := models.InsertFlightFacility(flightID, name)

	return newFlightFacility, nil
}
