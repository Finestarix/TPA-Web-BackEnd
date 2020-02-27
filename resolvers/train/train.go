package train

import (
	models "../../models/train"
	"errors"
	"github.com/graphql-go/graphql"
	"strconv"
	"time"
)

func AllTrain(p graphql.ResolveParams) (i interface{}, err error) {
	train := models.GetAllTrain()

	if len(train) == 0 {
		return nil, errors.New("error: no data found")
	}

	return train, nil
}

func InsertTrain(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	code := p.Args["code"].(string)
	arrival := p.Args["arrival"].(string)
	arrivalTime := p.Args["arrivalTime"].(string)
	transit := p.Args["transit"].(string)
	departure := p.Args["departure"].(string)
	departureTime := p.Args["departureTime"].(string)
	seat := p.Args["seat"].(int)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)

	newTrain := models.InsertTrain(name, code, arrival, arrivalTimeConv, transit, departure, departureTimeConv, seat, price)

	return newTrain, nil
}

func UpdateTrain(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(string)
	arrivalTime := p.Args["arrivalTime"].(string)
	departureTime := p.Args["departureTime"].(string)
	seat := p.Args["seat"].(int)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)
	idConv, _ := strconv.Atoi(id)

	newTrain := models.UpdateTrain(idConv, arrivalTimeConv, departureTimeConv, seat, price)

	return newTrain, nil
}

func DeleteTrain(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(string)

	idInt, _ := strconv.Atoi(id)

	train := models.DeleteTrain(idInt)

	return train, nil
}
