package flight

import (
models "../../models/flight"
"errors"
"github.com/graphql-go/graphql"
)

func AllFlightAirport(p graphql.ResolveParams) (i interface{}, err error) {
	flightAirport := models.GetAllFlightAirport()
	if len(flightAirport) == 0 {
		return nil, errors.New("error: no data found")
	}

	return flightAirport, nil
}

func InsertFlightAirport(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	code := p.Args["code"].(string)
	city := p.Args["city"].(string)

	newFlightAirport := models.InsertFlightAirport(name, code, city)

	return newFlightAirport, nil
}
