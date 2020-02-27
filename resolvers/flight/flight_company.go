package flight

import (
	models "../../models/flight"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllFlightCompany(p graphql.ResolveParams) (i interface{}, err error) {
	flightCompany := models.GetAllFlightCompany()
	if len(flightCompany) == 0 {
		return nil, errors.New("error: no data found")
	}

	return flightCompany, nil
}

func InsertFlightCompany(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	image := p.Args["image"].(string)

	newFlightCompany := models.InsertFlightCompany(name, image)

	return newFlightCompany, nil
}
