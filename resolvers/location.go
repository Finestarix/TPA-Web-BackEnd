package resolvers

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllLocation(p graphql.ResolveParams) (i interface{}, err error) {
	locations := models.GetAllLocation()

	if len(locations) == 0 {
		return nil, errors.New("error: no data found")
	}

	return locations, nil
}

func InsertLocation(p graphql.ResolveParams) (i interface{}, err error) {
	city := p.Args["province"].(string)
	province := p.Args["region"].(string)
	region := p.Args["region"].(string)

	newLocation := models.InsertLocation(city, province, region)

	return newLocation, nil
}
