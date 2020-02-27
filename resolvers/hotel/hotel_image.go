package hotel

import (
	models "../../models/hotel"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllHotelImage(p graphql.ResolveParams) (i interface{}, err error) {
	hotelImages := models.GetAllHotelImage()

	if len(hotelImages) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotelImages, nil
}

func InsertHotelImage(p graphql.ResolveParams) (i interface{}, err error) {
	hotelid := p.Args["hotelid"].(int)
	source := p.Args["source"].(string)

	newHotelImage := models.InsertHotelImage(hotelid, source)

	return newHotelImage, nil
}

