package hotel

import (
	models "../../models/hotel"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllHotelType(p graphql.ResolveParams) (i interface{}, err error) {
	hotelType := models.GetAllHotelType()

	if len(hotelType) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotelType, nil
}

func InsertHotelType(p graphql.ResolveParams) (i interface{}, err error) {
	hotelid := p.Args["hotelid"].(int)
	name := p.Args["name"].(string)

	newHotelType := models.InsertHotelType(hotelid, name)

	return newHotelType, nil
}
