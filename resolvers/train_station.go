package resolvers

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllTrainStation(p graphql.ResolveParams) (i interface{}, err error) {
	trainStation := models.GetAllTrainStation()

	if len(trainStation) == 0 {
		return nil, errors.New("error: no data found")
	}

	return trainStation, nil
}

func InsertTrainStation(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	city := p.Args["city"].(string)
	code := p.Args["code"].(string)

	newTrainStation := models.InsertTrainStation(name, code, city)

	return newTrainStation, nil
}
