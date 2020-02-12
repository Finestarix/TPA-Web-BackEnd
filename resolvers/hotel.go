package resolvers

import (
	"../models"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllHotel(p graphql.ResolveParams) (i interface{}, err error) {
	hotels := models.GetAllHotel()

	if len(hotels) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotels, nil
}

func InsertHotel(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	address := p.Args["address"].(string)
	location := p.Args["location"].(string)
	price := p.Args["price"].(int)
	rating := p.Args["rating"].(float64)
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)

	newHotel := models.InsertHotel(name, address, location, price, rating, latitude, longitude)

	return newHotel, nil
}

func GetNearestHotel(p graphql.ResolveParams) (i interface{}, err error) {
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)

	hotels := models.GetNearestHotel(latitude, longitude)

	if len(hotels) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotels, nil
}

func GetHotelByCity(p graphql.ResolveParams) (i interface{}, err error) {
	city := p.Args["city"].(string)

	hotels := models.GetHotelByCity(city)

	if len(hotels) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotels, nil
}

func GetHotelByProvince(p graphql.ResolveParams) (i interface{}, err error) {
	province := p.Args["province"].(string)

	hotels := models.GetHotelByProvince(province)

	if len(hotels) == 0 {
		return nil, errors.New("error: no data found")
	}

	return hotels, nil
}