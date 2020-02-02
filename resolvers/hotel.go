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
	city := p.Args["city"].(string)
	price := p.Args["price"].(int)
	rating := p.Args["rating"].(float64)

	newHotel := models.InsertHotel(name, address, city, price, rating)

	return newHotel, nil
}