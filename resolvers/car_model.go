package resolvers

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllCarModel(p graphql.ResolveParams) (i interface{}, err error) {
	cars := models.GetAllCarModel()

	if len(cars) == 0 {
		return nil, errors.New("error: no data found")
	}

	return cars, nil
}

func InsertCarModel(p graphql.ResolveParams) (i interface{}, err error) {
	brand := p.Args["brand"].(string)
	model := p.Args["model"].(string)
	passenger := p.Args["passenger"].(int)
	baggage := p.Args["baggage"].(int)
	image := p.Args["image"].(string)

	newCars := models.InsertCarModel(brand, model, passenger, baggage, image);

	return newCars, nil
}

