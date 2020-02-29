package flight

import (
	models "../../models/flight"
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

func AllFlight(p graphql.ResolveParams) (i interface{}, err error) {
	flight := models.GetAllFlight()
	if len(flight) == 0 {
		return nil, errors.New("error: no data found")
	}

	return flight, nil
}

func InsertFlight(p graphql.ResolveParams) (i interface{}, err error) {
	companyName := p.Args["company"].(string)
	fromAirportName := p.Args["fromAirport"].(string)
	toAirportName := p.Args["toAirport"].(string)
	transitAirportName := p.Args["transitAirport"].(string)
	arrivalTime := p.Args["arrivalTime"].(string)
	departureTime := p.Args["departureTime"].(string)
	model := p.Args["model"].(string)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)

	newFlight := models.InsertFlight(companyName, fromAirportName, arrivalTimeConv, toAirportName, departureTimeConv, price, model, transitAirportName)

	return newFlight, nil
}

func UpdateFlight(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	fromAirportName := p.Args["fromAirport"].(string)
	toAirportName := p.Args["toAirport"].(string)
	transitAirportName := p.Args["transitAirport"].(string)
	arrivalTime := p.Args["arrivalTime"].(string)
	departureTime := p.Args["departureTime"].(string)
	model := p.Args["model"].(string)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)

	flight := models.UpdateFlight(id, fromAirportName, arrivalTimeConv, toAirportName, departureTimeConv, price, model, transitAirportName)

	return flight, nil
}

func DeleteFlight(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	flight := models.DeleteFlight(id)

	return flight, nil
}
