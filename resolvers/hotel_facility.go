package resolvers

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllHotelFacility(p graphql.ResolveParams) (i interface{}, err error) {
	hotelFacility := models.GetAllHotelFacility()

	if len(hotelFacility) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotelFacility, nil
}

func InsertHotelFacility(p graphql.ResolveParams) (i interface{}, err error) {
	hotelid := p.Args["hotelid"].(int)
	name := p.Args["name"].(string)

	newHotelFacility := models.InsertHotelFacility(hotelid, name)

	return newHotelFacility, nil
}
